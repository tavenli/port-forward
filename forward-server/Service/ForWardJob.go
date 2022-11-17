package Service

import (
	"fmt"
	"forward-core/Constant"
	"forward-core/Models"
	"forward-core/NetUtils"
	"forward-core/Utils"
	"io"
	"net"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
)

type ForWardJob struct {
	Config        *Models.ForwardConfig
	ClientMap     map[string]*ForWardClient
	ClientMapLock sync.Mutex
	Status        byte
	PortListener  net.Listener
	UdpForwardJob *UdpForward
}

func (_self *ForWardJob) StartJob(result chan Models.FuncResult) {

	sourceAddr := fmt.Sprint(_self.Config.SrcAddr, ":", _self.Config.SrcPort)
	destAddr := fmt.Sprint(_self.Config.DestAddr, ":", _self.Config.DestPort)

	resultData := &Models.FuncResult{Code: 0, Msg: "success"}
	var err error
	if _self.IsUdpJob() {
		//_self.PortListener, err = NetUtils.NewKCP(sourceAddr, Common.DefaultKcpSetting())
		//_self.UdpForwardJob.UdpListenerConn, err = NetUtils.NewUDP(sourceAddr)

		err = _self.UdpForwardJob.DoUdpForward(sourceAddr, destAddr)

		if err != nil {
			logs.Error("启动UDP监听 ", sourceAddr, " 出错：", err)
			resultData.Code = 1
			resultData.Msg = fmt.Sprint("启动UDP监听 ", sourceAddr, " 出错：", err)
			result <- *resultData
			return
		}

		_self.Status = Constant.RunStatus_Running
		logs.Debug("启动UDP端口转发，从 ", sourceAddr, " 到 ", destAddr)
		result <- *resultData

	} else {
		_self.PortListener, err = NetUtils.NewTCP(sourceAddr)

		if err != nil {
			logs.Error("启动监听 ", sourceAddr, " 出错：", err)
			resultData.Code = 1
			resultData.Msg = fmt.Sprint("启动监听 ", sourceAddr, " 出错：", err)
			result <- *resultData
			return
		}

		_self.Status = Constant.RunStatus_Running
		logs.Debug("启动端口转发，从 ", sourceAddr, " 到 ", destAddr)
		result <- *resultData

		_self.doTcpForward(destAddr)

	}

}

func (_self *ForWardJob) doTcpForward(destAddr string) {

	for {
		realClientConn, err := _self.PortListener.Accept()
		if err != nil {
			logs.Error("Forward Accept err:", err.Error())
			logs.Error(fmt.Sprint("转发出现异常：", _self.Config.SrcAddr, ":", _self.Config.SrcPort, "->", destAddr))
			_self.StopJob()
			break
		}

		if ForWardDebug == true {
			logs.Info("新用户 ", realClientConn.RemoteAddr().String(), " 数据转发规则：", fmt.Sprint(_self.Config.SrcAddr, ":", _self.Config.SrcPort), "->", destAddr)
		}

		var destConn net.Conn
		if _self.Config.Protocol == "UDP" {
			//destConn, err = Common.DialKcpTimeout(destAddr, 100)
			destConn, err = net.DialTimeout("UDP", destAddr, 30*time.Second)
		} else {
			destConn, err = net.DialTimeout("tcp", destAddr, 30*time.Second)
		}

		if err != nil {
			if ForWardDebug == true {
				logs.Warn("转发出现异常 Forward to Dest Addr err:", err.Error())
			}

			//break
			continue

		}

		forwardClient := &ForWardClient{realClientConn, destConn, nil, _self.ClosedCallBack}

		if Utils.IsNotEmpty(_self.Config.Others) {
			var dispatchConns []io.Writer
			//分发方式
			dispatchTargets := Utils.Split(_self.Config.Others, ";")

			for _, dispatchTarget := range dispatchTargets {
				logs.Debug("分发到：", dispatchTarget)
				dispatchTargetConn, err := net.DialTimeout("tcp", dispatchTarget, 30*time.Second)
				if err == nil {
					dispatchConns = append(dispatchConns, dispatchTargetConn)
				}

			}

			forwardClient.DispatchConns = dispatchConns

			go forwardClient.DispatchData(dispatchConns)
		} else {
			go forwardClient.StartForward()
		}

		_self.RegistryClient(_self.GetClientId(realClientConn), forwardClient)
		//_self.RegistryClient(fmt.Sprint(sourceAddr, "_", "TCP", "_", id), forwardClient)

	}
}

func (_self *ForWardJob) ClosedCallBack(srcConn net.Conn, destConn net.Conn) {

	_self.UnRegistryClient(_self.GetClientId(srcConn))
}

func (_self *ForWardJob) GetClientId(conn net.Conn) string {
	return conn.RemoteAddr().String()
}

func (_self *ForWardJob) RegistryClient(srcAddr string, forwardClient *ForWardClient) {
	_self.ClientMapLock.Lock()
	defer _self.ClientMapLock.Unlock()

	_self.ClientMap[srcAddr] = forwardClient

}

func (_self *ForWardJob) UnRegistryClient(srcAddr string) {
	_self.ClientMapLock.Lock()
	defer _self.ClientMapLock.Unlock()

	delete(_self.ClientMap, srcAddr)
	if ForWardDebug == true {
		logs.Debug("UnRegistryClient srcAddr: ", srcAddr)
	}

}

func (_self *ForWardJob) IsJobRunning() bool {

	return _self.Status == Constant.RunStatus_Running

}

func (_self *ForWardJob) IsUdpJob() bool {
	return Utils.ToUpper(_self.Config.Protocol) == "UDP"
}

func (_self *ForWardJob) StopJob() {

	if _self.IsUdpJob() {
		_self.stopUdpJob()
	} else {
		_self.stopTcpJob()
	}

	_self.Status = Constant.RunStatus_Stoped
}

func (_self *ForWardJob) stopTcpJob() {

	_self.PortListener.Close()

	for srcAddr, client := range _self.ClientMap {
		if ForWardDebug == true {
			logs.Debug("停止真实用户连接：", srcAddr)
		}
		client.StopForward()
	}

	_self.ClientMap = nil
}

func (_self *ForWardJob) stopUdpJob() {

	_self.UdpForwardJob.Close()
}
