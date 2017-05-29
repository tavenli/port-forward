package services

import (
	"io"
	"net"
	"port-forward/models"
	"strings"
	"sync"
	"time"

	"fmt"

	"github.com/astaxie/beego/logs"
)

type ForwardService struct {
}

var (
	portMap       = make(map[string]net.Listener)
	portMapLock   = new(sync.Mutex)
	clientMap     = make(map[string]net.Conn)
	clientMapLock = new(sync.Mutex)
)

func init() {

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

func (_self *ForwardService) GetKeyByEntity(entity *models.PortForward) string {

	fromAddr := fmt.Sprint(entity.Addr, ":", entity.Port)
	toAddr := fmt.Sprint(entity.TargetAddr, ":", entity.TargetPort)
	key := _self.GetKey(fromAddr, toAddr)

	return key
}

func (_self *ForwardService) GetKey(sourcePort, targetPort string) string {

	return fmt.Sprint(sourcePort, "_TCP_", targetPort)

}

//
// sourcePort 源地址和端口，例如：0.0.0.0:8700，本程序会新建立监听
// targetPort 数据转发给哪个端口，例如：192.168.1.100:3306
func (_self *ForwardService) StartPortForward(sourcePort string, targetPort string, result chan models.ResultData) {
	resultData := &models.ResultData{Code: 0, Msg: ""}
	logs.Debug("StartTcpPortForward sourcePort: ", sourcePort, " targetPort:", targetPort)

	key := _self.GetKey(sourcePort, targetPort)

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
		_self.RegistryClient(fmt.Sprint(sourcePort, "_", id), sourceConn)

		logs.Debug("conn.RemoteAddr().String() ：", id)

		//targetPort := "172.16.128.83:22"
		targetConn, err := net.DialTimeout("tcp", targetPort, 30*time.Second)

		go func() {
			_, err = _self.Copy(targetConn, sourceConn)
			if err != nil {
				logs.Error("1网络连接异常：", err)
				_self.UnRegistryClient(fmt.Sprint(sourcePort, "_", sourceConn.RemoteAddr().String()))
			}
		}()

		go func() {
			_, err = _self.Copy(sourceConn, targetConn)
			if err != nil {
				logs.Error("2网络连接异常：", err)
				_self.UnRegistryPort(sourcePort)
			}
		}()

	}

	logs.Debug("TcpPortForward sourcePort: ", sourcePort, " Close.")

}

func (_self *ForwardService) ClosePortForward(sourcePort string, targetPort string, result chan models.ResultData) {
	resultData := &models.ResultData{Code: 0, Msg: ""}

	logs.Debug("CloseTcpPortForward:", sourcePort)
	//先关闭客户端连接
	for cId, conn := range clientMap {
		//logs.Debug("clientMap id：", cId)
		if strings.HasPrefix(cId, sourcePort+"_") {
			logs.Debug("close clientMap id：", cId)
			conn.Close()
			_self.UnRegistryClient(cId)
		}

	}

	//关闭本地监听
	key := _self.GetKey(sourcePort, targetPort)
	if localListener, ok := portMap[key]; ok {
		localListener.Close()
		logs.Debug("listener close:", key)
		_self.UnRegistryPort(key)
	} else {
		resultData.Code = 1
		resultData.Msg = fmt.Sprint("未启用监听 ", key)

	}

	result <- *resultData

	logs.Debug("CloseTcpPortForward finished.")

}

func (_self *ForwardService) Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}
