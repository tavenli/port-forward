package Service

import (
	"fmt"
	"forward-core/Constant"
	"forward-core/Models"
	"net"
	"sync"

	"github.com/astaxie/beego/logs"
)

type ForWardServer struct {
	JobMap     map[string]*ForWardJob
	JobMapLock sync.Mutex
}

func NewForWardServer() *ForWardServer {
	return &ForWardServer{
		JobMap: make(map[string]*ForWardJob, 200),
	}
}

func (_self *ForWardServer) FindAllForward() []*Models.ForwardInfo {
	var forwardList []*Models.ForwardInfo
	for _, forWardJob := range _self.JobMap {

		forwardInfo := new(Models.ForwardInfo)

		forwardInfo.Name = forWardJob.Config.Name
		forwardInfo.Status = forWardJob.Status
		forwardInfo.Protocol = forWardJob.Config.Protocol
		forwardInfo.SrcAddr = forWardJob.Config.SrcAddr
		forwardInfo.SrcPort = forWardJob.Config.SrcPort
		forwardInfo.DestAddr = forWardJob.Config.DestAddr
		forwardInfo.DestPort = forWardJob.Config.DestPort

		if forWardJob.IsUdpJob() {
			for key, _ := range forWardJob.UdpForwardJob.UdpConns {

				forwardInfo.Clients = append(forwardInfo.Clients, key)
			}
		} else {
			for _, client := range forWardJob.ClientMap {
				forwardInfo.Clients = append(forwardInfo.Clients, client.SrcConn.RemoteAddr().String())
			}
		}

		forwardInfo.OnlineCount = len(forwardInfo.Clients)

		forwardList = append(forwardList, forwardInfo)
	}

	return forwardList
}

func (_self *ForWardServer) GetForwardInfo(config *Models.ForwardConfig) *Models.ForwardInfo {

	forwardInfo := new(Models.ForwardInfo)
	forWardJob := _self.GetRegistryJob(config)
	if forWardJob != nil {
		forwardInfo.Name = forWardJob.Config.Name
		forwardInfo.Status = forWardJob.Status
		forwardInfo.Protocol = forWardJob.Config.Protocol
		forwardInfo.SrcAddr = forWardJob.Config.SrcAddr
		forwardInfo.SrcPort = forWardJob.Config.SrcPort
		forwardInfo.DestAddr = forWardJob.Config.DestAddr
		forwardInfo.DestPort = forWardJob.Config.DestPort

		for _, client := range forWardJob.ClientMap {
			forwardInfo.Clients = append(forwardInfo.Clients, client.SrcConn.RemoteAddr().String())
		}

		forwardInfo.OnlineCount = len(forwardInfo.Clients)

	}

	return forwardInfo
}

func (_self *ForWardServer) GetForwardJob(config *Models.ForwardConfig) *ForWardJob {
	return _self.GetRegistryJob(config)
}

func (_self *ForWardServer) OpenForward(config *Models.ForwardConfig, result chan Models.FuncResult) {
	hasJob := _self.GetForwardJob(config)
	if hasJob != nil && hasJob.Status == Constant.RunStatus_Running {
		resultData := &Models.FuncResult{Code: 1, Msg: "该端口转发正在执行中"}
		result <- *resultData
		return
	}

	forWardJob := new(ForWardJob)
	forWardJob.ClientMap = make(map[string]*ForWardClient, 500)
	forWardJob.Config = config
	forWardJob.UdpForwardJob = NewUdpForward()

	go forWardJob.StartJob(result)

	_self.RegistryJob(config, forWardJob)

}

func (_self *ForWardServer) GetJobKey(config *Models.ForwardConfig) string {
	srcAddr := fmt.Sprint(config.SrcAddr, ":", config.SrcPort)
	destAddr := fmt.Sprint(config.DestAddr, ":", config.DestPort)

	return fmt.Sprint(srcAddr, "_", config.Protocol, "_", destAddr)
}

func (_self *ForWardServer) GetClientId(conn net.Conn) string {
	return conn.RemoteAddr().String()
}

func (_self *ForWardServer) RegistryJob(config *Models.ForwardConfig, forWardJob *ForWardJob) {
	_self.JobMapLock.Lock()
	defer _self.JobMapLock.Unlock()

	_self.JobMap[_self.GetJobKey(config)] = forWardJob

}

func (_self *ForWardServer) UnRegistryJob(config *Models.ForwardConfig) {
	_self.JobMapLock.Lock()
	defer _self.JobMapLock.Unlock()

	key := _self.GetJobKey(config)
	delete(_self.JobMap, key)
	if ForWardDebug == true {
		logs.Debug("UnRegistryClient key: ", key)
	}

}

func (_self *ForWardServer) GetRegistryJob(config *Models.ForwardConfig) *ForWardJob {
	if forWardJob, ok := _self.JobMap[_self.GetJobKey(config)]; ok {
		return forWardJob
	}

	return nil
}

func (_self *ForWardServer) CloseForward(config *Models.ForwardConfig) {

	forWardJob := _self.GetRegistryJob(config)
	if forWardJob != nil {
		logs.Debug("停止转发，找到执行者：", _self.GetJobKey(config))
		forWardJob.StopJob()
		_self.UnRegistryJob(config)
	}

}

func (_self *ForWardServer) CloseAllForward() {

	for _, forWardJob := range _self.JobMap {
		forWardJob.StopJob()
		//delete(_self.JobMap, key)
		_self.UnRegistryJob(forWardJob.Config)

	}

	//_self.JobMap = nil
	_self.JobMap = make(map[string]*ForWardJob, 200)

}
