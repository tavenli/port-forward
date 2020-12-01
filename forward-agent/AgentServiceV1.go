package main

import (
	"bufio"
	"forward-core/Constant"
	"forward-core/NetUtils"
	"github.com/astaxie/beego/logs"
	"io"
	"net"
	"strconv"
	"time"
)

type AgentServiceV1 struct {
	localConnMap map[int]net.Conn
	MagicServerAddr string
	AgentOnline     bool
}


func (_self *AgentServiceV1) ConnToMagicServer() {
	serviceConn, err := net.DialTimeout("tcp", _self.MagicServerAddr, 30*time.Second)

	if err != nil {
		logs.Error("try dial err", err)
		_self.AgentOnline = false
		return
	}

	callback := func(conn net.Conn, sessionId int, cmd byte, payload []byte) {
		//payload 收到的消息内容
		_self.OnTunnelRecv(conn, sessionId, cmd, payload)

	}
	logs.Debug("开始接收服务端返回指令或数据...")
	_self.AgentOnline = true
	go NetUtils.ReadConn(serviceConn, callback)
}

func (_self *AgentServiceV1) OnTunnelRecv(conn net.Conn, sessionId int, cmd byte, payload []byte) {
	logs.Debug("收到一条给 sessionId：", sessionId, " 客户端的数据，指令是：", cmd)

	switch cmd {
	case Constant.MagicCmd_AgentListenerOpen:
		targetAddr := string(payload)
		go _self.ListenForClient(targetAddr, _self.MagicServerAddr)
	case Constant.MagicCmd_AgentConnOpen:
		targetAddr := string(payload)
		logs.Debug("sessionId：", sessionId, " 收到 AgentConnOpen 指令是，打开本地连接：", targetAddr)
		//AgentConnOpen 让连接进来的客户端，在它的本地创建一个连接，并关联好sessionId
		localConn, err := net.DialTimeout("tcp", targetAddr, 30*time.Second)
		if err != nil {
			logs.Error("try dial err", err)
			return
		}
		_self.localConnMap[sessionId] = localConn
		//接收 localConn 返回数据，并将返回的数据，写回给 conn，带上 sessionId
		go _self.ReadRawConn(localConn, conn, sessionId, Constant.MagicCmd_DataToMagic)

	case Constant.MagicCmd_DataToAgent:
		logs.Debug("sessionId：", sessionId, " 收到 MsgToAgent 指令")
		localConn := _self.localConnMap[sessionId]
		localConn.Write(payload)
		logs.Debug("sessionId：", sessionId, " 数据已写入本地目标连接")
	case Constant.MagicCmd_Refused:
		//client := string(payload)
		logs.Debug("Magic服务端拒绝本次连接")
	}

}

func (_self *AgentServiceV1) ReadRawConn(from net.Conn, magic_client_Conn net.Conn, sessionId int, cmd byte) {

	arr := make([]byte, 5000)
	reader := bufio.NewReader(from)

	for {
		size, err := reader.Read(arr)
		if err != nil {
			break
		}

		err = NetUtils.WriteConn(magic_client_Conn, sessionId, cmd, arr[0:size])

		if err != nil {
			//有异常
			logs.Error(err)
			break
		}
	}
}

func (_self *AgentServiceV1) ListenForClient(localListenAddr, toAddr string) {
	client_listener, err := net.Listen("tcp", localListenAddr)
	if err != nil {
		logs.Error("ListenForClient err:", err)
		return
	}

	for {
		logs.Debug("ListenForClient Ready to Accept ...")
		client_Conn, err := client_listener.Accept()
		if err != nil {
			logs.Error("Accept err:", err)
			break
		}

		//连接到远程服务
		serviceConn, err := net.DialTimeout("tcp", toAddr, 30*time.Second)

		if err != nil {
			logs.Error("try dial err", err)
			return
		}

		go func() {
			_, err = io.Copy(serviceConn, client_Conn)
			if err != nil {
				logs.Error("to magic_client 网络连接异常：", err)
			}
		}()

		go func() {
			_, err = io.Copy(client_Conn, serviceConn)
			if err != nil {
				logs.Error("to magic_client 网络连接异常2：", err)
			}
		}()
	}

}

func (_self *AgentServiceV1) ConnectToMagicLoop() {
	//客户端与服务端建立连接
	if !_self.AgentOnline {
		_self.ConnToMagicServer()
		delay := 3
		time.AfterFunc(time.Second*time.Duration(delay), func() {
			_self.ConnectToMagicLoop()
		})
		logs.Debug("reConnect after " + strconv.Itoa(delay) + " seconds")
	}

}
