package services

import (
	"bufio"
	"forward-core/common"
	"io"
	"net"
	"port-forward/models"
	"port-forward/utils"
	"strings"
	"sync"
	"time"

	"fmt"

	"github.com/astaxie/beego/logs"
)

type ForwardService struct {
}

var (
	portMap            = make(map[string]net.Listener)
	portMapLock        = new(sync.Mutex)
	clientMap          = make(map[string]net.Conn)
	clientMapLock      = new(sync.Mutex)
	magicClientMap     = make(map[string]net.Conn)
	magicClientMapLock = new(sync.Mutex)
	magicListener      net.Listener
	sessionId          = 0
	idLock             = new(sync.Mutex)
	sessionConnMap     = make(map[int]net.Conn)
	agentRunType       = 0
	magicTargetAddr    = ""
)

func init() {

}

func (_self *ForwardService) GetNewSessionId() int {
	idLock.Lock()
	defer idLock.Unlock()
	sessionId++

	return sessionId
}

func (_self *ForwardService) PortConflict(key string) bool {
	portMapLock.Lock()
	defer portMapLock.Unlock()

	if _, ok := portMap[key]; ok {
		return true
	} else {
		return false
	}

}

func (_self *ForwardService) RegistryPort(key string, listener net.Listener) {
	portMapLock.Lock()
	defer portMapLock.Unlock()

	portMap[key] = listener

}

func (_self *ForwardService) UnRegistryPort(key string) {
	portMapLock.Lock()
	defer portMapLock.Unlock()

	delete(portMap, key)
	logs.Debug("UnRegistryPort key: ", key)

}

func (_self *ForwardService) RegistryClient(sourcePort string, conn net.Conn) {
	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	clientMap[sourcePort] = conn

}

func (_self *ForwardService) UnRegistryClient(sourcePort string) {
	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	delete(clientMap, sourcePort)
	logs.Debug("UnRegistryClient sourcePort: ", sourcePort)

}

func (_self *ForwardService) RegistryMagicClient(key string, conn net.Conn) {
	magicClientMapLock.Lock()
	defer magicClientMapLock.Unlock()

	magicClientMap[key] = conn

}

func (_self *ForwardService) GetTopMagicClient() net.Conn {
	magicClientMapLock.Lock()
	defer magicClientMapLock.Unlock()

	for _, v := range magicClientMap {
		return v
	}

	return nil

}

func (_self *ForwardService) GetMagicListener() net.Listener {

	return magicListener

}

func (_self *ForwardService) UnRegistryMagicClient(key string) {
	magicClientMapLock.Lock()
	defer magicClientMapLock.Unlock()

	delete(magicClientMap, key)
	logs.Debug("UnRegistryMagicClient key: ", key)

}

func (_self *ForwardService) CountMagicClient() int {
	magicClientMapLock.Lock()
	defer magicClientMapLock.Unlock()

	return len(magicClientMap)

}

func (_self *ForwardService) GetMagicClient() map[string]net.Conn {

	return magicClientMap

}

func (_self *ForwardService) GetKeyByEntity(entity *models.PortForward) string {

	fromAddr := fmt.Sprint(entity.Addr, ":", entity.Port)
	toAddr := fmt.Sprint(entity.TargetAddr, ":", entity.TargetPort)
	key := _self.GetKey(fromAddr, toAddr, entity.FType)

	return key
}

func (_self *ForwardService) GetKey(sourcePort, targetPort string, fType int) string {

	return fmt.Sprint(sourcePort, "_", fType, "_TCP_", targetPort)

}

func (_self *ForwardService) StartPortForward(portForward *models.PortForward, result chan models.ResultData) {
	if portForward.FType == 0 {
		_self.StartPortToPortForward(portForward, result)
	} else {
		_self.StartMagicForward(portForward, result)
	}
}

//
// sourcePort 源地址和端口，例如：0.0.0.0:8700，本程序会新建立监听
// targetPort 数据转发给哪个端口，例如：192.168.1.100:3306
func (_self *ForwardService) StartPortToPortForward(portForward *models.PortForward, result chan models.ResultData) {

	sourcePort := fmt.Sprint(portForward.Addr, ":", portForward.Port)
	targetPort := fmt.Sprint(portForward.TargetAddr, ":", portForward.TargetPort)
	fType := portForward.FType

	resultData := &models.ResultData{Code: 0, Msg: ""}
	logs.Debug("StartTcpPortForward sourcePort: ", sourcePort, " targetPort:", targetPort)

	key := _self.GetKey(sourcePort, targetPort, fType)

	if _self.PortConflict(key) {
		resultData.Code = 1
		resultData.Msg = fmt.Sprint("监听地址已被占用 ", sourcePort)
		result <- *resultData
		return
	}

	localListener, err := net.Listen("tcp", sourcePort)

	if err != nil {
		logs.Error("启动监听 ", sourcePort, " 出错：", err)
		resultData.Code = 1
		resultData.Msg = fmt.Sprint("启动监听 ", sourcePort, " 出错：", err)
		result <- *resultData
		return
	}

	_self.RegistryPort(key, localListener)

	result <- *resultData

	for {
		logs.Debug("Ready to Accept ...")
		sourceConn, err := localListener.Accept()

		if err != nil {
			logs.Error("Accept err:", err)
			break
		}

		id := sourceConn.RemoteAddr().String()
		_self.RegistryClient(fmt.Sprint(sourcePort, "_", fType, "_", id), sourceConn)

		logs.Debug("conn.RemoteAddr().String() ：", id)

		//targetPort := "172.16.128.83:22"
		targetConn, err := net.DialTimeout("tcp", targetPort, 30*time.Second)

		if utils.IsNotEmpty(portForward.Others) {
			var dispatchConns []io.Writer
			dispatchConns = append(dispatchConns, targetConn)
			//分发方式
			dispatchTargets := utils.Split(portForward.Others, ";")

			for _, dispatchTarget := range dispatchTargets {
				logs.Debug("分发到：", dispatchTarget)
				dispatchTargetConn, err := net.DialTimeout("tcp", dispatchTarget, 30*time.Second)
				if err == nil {
					dispatchConns = append(dispatchConns, dispatchTargetConn)
				}

			}

			go func() {
				mWriter := io.MultiWriter(dispatchConns...)
				_, err = _self.Copy(mWriter, sourceConn)
				if err != nil {
					logs.Error("Dispatch网络连接异常：", err)
				}
			}()

		} else {
			go func() {
				_, err = _self.Copy(targetConn, sourceConn)
				if err != nil {
					logs.Error("客户端来源数据转发到目标端口异常：", err)
					_self.UnRegistryClient(fmt.Sprint(sourcePort, "_", fType, "_", sourceConn.RemoteAddr().String()))
				}
			}()
		}

		go func() {
			_, err = _self.Copy(sourceConn, targetConn)
			if err != nil {
				logs.Error("目标端口返回响应数据异常：", err)
				_self.UnRegistryPort(key)
			}
		}()

	}

	logs.Debug("TcpPortForward sourcePort: ", sourcePort, " Close.")

}

func (_self *ForwardService) DataDispatch(src io.Reader, targetPorts []string) {
	for _, target := range targetPorts {
		logs.Debug("分发到：", target)
		go func() {
			targetConn, err := net.DialTimeout("tcp", target, 30*time.Second)
			_, err = _self.Copy(targetConn, src)
			if err != nil {
				logs.Error("Dispatch网络连接异常：", err)
			}
		}()
	}

}

func (_self *ForwardService) ClosePortForward(sourcePort string, targetPort string, fType int, result chan models.ResultData) {
	resultData := &models.ResultData{Code: 0, Msg: ""}

	logs.Debug("CloseTcpPortForward:", sourcePort)
	//先关闭客户端连接
	for cId, conn := range clientMap {
		//logs.Debug("clientMap id：", cId)
		if strings.HasPrefix(cId, fmt.Sprint(sourcePort, "_", fType)) {
			logs.Debug("close clientMap id：", cId)
			if conn != nil {
				conn.Close()
			}
			_self.UnRegistryClient(cId)
		}

	}

	//关闭本地监听
	key := _self.GetKey(sourcePort, targetPort, fType)
	if localListener, ok := portMap[key]; ok {
		if localListener != nil {
			localListener.Close()
			logs.Debug("listener close:", key)
		}

		_self.UnRegistryPort(key)
	} else {
		resultData.Code = 1
		resultData.Msg = fmt.Sprint("未启用监听 ", key)

	}

	if fType == 1 {
		agentRunType = 0
	}

	result <- *resultData

	logs.Debug("CloseTcpPortForward finished.")

}

func (_self *ForwardService) Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}

func (_self *ForwardService) StartMagicService(addr string, result chan models.ResultData) {
	//启动穿透服务端
	resultData := &models.ResultData{Code: 0, Msg: ""}
	var err error
	magicListener, err = net.Listen("tcp", addr)
	if err != nil {
		logs.Error("Magic Listen err:", err)
		resultData.Code = 1
		resultData.Msg = err.Error()
		result <- *resultData
		return
	}

	result <- *resultData

	for {
		logs.Debug("Magic Ready to Accept ...")
		magic_client_Conn, err := magicListener.Accept()
		if err != nil {
			logs.Error("Accept err:", err)
			break
		}

		if _self.CountMagicClient() > 0 && _self.CurrentAgentRunType() != 1 {
			logs.Debug("目前版本只支持一个Agent连接，后续会增加多个的支持")
			NetCommonS.WriteConn(magic_client_Conn, -1, common.MagicRefused, []byte(""))
			magic_client_Conn.Close()
			continue
		}

		if _self.CountMagicClient() == 0 {
			magicId := magic_client_Conn.RemoteAddr().String()
			_self.RegistryMagicClient(magicId, magic_client_Conn)
		} else {
			if _self.CurrentAgentRunType() == 1 {
				_self.MagicJustCopy(magic_client_Conn, magicTargetAddr)
			}

		}

	}

}

func (_self *ForwardService) StopMagicService(result chan models.ResultData) {
	resultData := &models.ResultData{Code: 0, Msg: ""}

	for k, conn := range magicClientMap {
		conn.Close()
		_self.UnRegistryMagicClient(k)

	}
	magicListener.Close()
	magicListener = nil

	result <- *resultData

}

func (_self *ForwardService) StartMagicForward(portForward *models.PortForward, result chan models.ResultData) {
	resultData := &models.ResultData{Code: 0, Msg: ""}

	agentConn := _self.GetTopMagicClient()

	if agentConn == nil {
		resultData.Code = 1
		resultData.Msg = "未检测到Agent连接"
		result <- *resultData
		return
	}

	if _self.CurrentAgentRunType() != 0 {
		resultData.Code = 1
		resultData.Msg = "有正在执行的Agent连接，开启转发失败"
		result <- *resultData
		return
	}

	if portForward.FType == 2 {
		//执行反向映射
		go _self.ReverseListenForClient(portForward, agentConn, result)
		callback := func(conn net.Conn, sessionId int, cmd byte, payload []byte) {
			//payload 收到的消息内容
			_self.OnTunnelRecv(sessionConnMap[sessionId], sessionId, cmd, payload)
		}
		logs.Debug("从 magic_client_Conn 读，写入到 client_Conn")
		go NetCommonS.ReadConn(agentConn, callback)
	} else {
		//发送指令
		localListenAddr := fmt.Sprint(portForward.Addr, ":", portForward.Port)
		NetCommonS.WriteConn(agentConn, -1, common.AgentListenerOpen, []byte(localListenAddr))
		result <- *resultData
		agentRunType = 1
		magicTargetAddr = fmt.Sprint(portForward.TargetAddr, ":", portForward.TargetPort)
		key := _self.GetKeyByEntity(portForward)
		_self.RegistryPort(key, nil)
	}

}

func (_self *ForwardService) MagicJustCopy(toConn net.Conn, targetAddr string) {

	localConn, err := net.DialTimeout("tcp", targetAddr, 30*time.Second)
	if err != nil {
		logs.Error("try dial err", err)
		return
	}

	go func() {
		_, err = io.Copy(localConn, toConn)
		if err != nil {
			logs.Error("JustCopy to local 网络连接异常：", err)
			localConn.Close()
		}
	}()

	go func() {
		_, err = io.Copy(toConn, localConn)
		if err != nil {
			logs.Error("JustCopy to local 网络连接异常2：", err)
			toConn.Close()
		}
	}()

}

func (_self *ForwardService) ReverseListenForClient(portForward *models.PortForward, magic_client_Conn net.Conn, result chan models.ResultData) {
	resultData := &models.ResultData{Code: 0, Msg: ""}

	localListenAddr := fmt.Sprint(portForward.Addr, ":", portForward.Port)
	//让客户端在本地建立连接与目标端口的连接
	remote := fmt.Sprint(portForward.TargetAddr, ":", portForward.TargetPort)
	fType := portForward.FType

	client_listener, err := net.Listen("tcp", localListenAddr)
	if err != nil {
		logs.Error("ListenForClient err:", err)
		resultData.Code = 1
		resultData.Msg = err.Error()
		result <- *resultData
		return
	}

	result <- *resultData
	agentRunType = 2
	key := _self.GetKeyByEntity(portForward)
	_self.RegistryPort(key, client_listener)

	//从 client_Conn 读，写入到 magic_client_Conn
	//从 magic_client_Conn 读，写入到 client_Conn
	for {
		logs.Debug("ListenForClient Ready to Accept ...")
		client_Conn, err := client_listener.Accept()
		if err != nil {
			logs.Error("Accept err:", err)
			break
		}

		id := client_Conn.RemoteAddr().String()
		_self.RegistryClient(fmt.Sprint(localListenAddr, "_", fType, "_", id), client_Conn)

		//有连接进来了，就创建一个sessionId
		sessionId := _self.GetNewSessionId()
		sessionConnMap[sessionId] = client_Conn
		logs.Debug("进来了一个连接，sessionId：", sessionId)

		NetCommonS.WriteConn(magic_client_Conn, sessionId, common.AgentConnOpen, []byte(remote))

		logs.Debug("向 sessionId：", sessionId, " 发送 AgentConnOpen 指令")

		logs.Debug("从 client_Conn 读，写入到 magic_client_Conn sessionId：", sessionId)
		go _self.ReadRawConn(client_Conn, magic_client_Conn, sessionId, common.MsgToAgent)

	}

}

func (_self *ForwardService) OnTunnelRecv(client_Conn net.Conn, sessionId int, cmd byte, payload []byte) {
	logs.Debug("收到一条给 sessionId：", sessionId, " 客户端的数据，指令是：", cmd)
	switch cmd {
	case common.MsgToMagic:
		client_Conn.Write(payload)
	}

}

func (_self *ForwardService) ReadRawConn(from net.Conn, magic_client_Conn net.Conn, sessionId int, cmd byte) {

	arr := make([]byte, 5000)
	reader := bufio.NewReader(from)

	for {
		size, err := reader.Read(arr)
		if err != nil {
			break
		}

		err = NetCommonS.WriteConn(magic_client_Conn, sessionId, cmd, arr[0:size])

		if err != nil {
			//有异常
			logs.Error(err)
			break
		}
	}
}

func (_self *ForwardService) CurrentAgentRunType() int {
	// 0:空闲，1：服务端映射到内网中，2：内网映射到服务端中
	return agentRunType
}
