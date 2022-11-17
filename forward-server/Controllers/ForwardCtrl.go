package Controllers

import (
	"fmt"
	"forward-core/Constant"
	"forward-core/Models"
	"forward-core/Utils"
	"forward-server/Controllers/BaseCtrl"
	"forward-server/Service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ForwardCtrl struct {
	BaseCtrl.ConsoleCtrl
}

// @router /u/ForwardList [get]
func (c *ForwardCtrl) ForwardList() {

	c.Data["ForWardDebug"] = Service.ForWardDebug

	c.TplName = "ucenter/forwardList.html"
}

// @router /u/ForwardList/json [post]
func (c *ForwardCtrl) ForwardListJson() {
	pageParam := new(Models.PageParam)
	pageParam.PIndex, _ = c.GetInt64("pIndex")
	pageParam.PSize, _ = c.GetInt64("pSize")

	port, _ := c.GetInt("port")
	targetAddr := c.GetString("targetAddr", "")
	targetPort, _ := c.GetInt("targetPort")

	query := &Models.PortForward{}
	query.Port = port
	query.TargetAddr = targetAddr
	query.TargetPort = targetPort
	query.FType = -1

	pageData := Service.SysDataS.GetPortForwardList(query, pageParam.PIndex, pageParam.PSize)

	for _, entity := range pageData.Data.([]*Models.PortForward) {
		forwardJob := Service.SysDataS.GetForwardJob(entity)
		//entity.Status = Utils.If(forwardJob!=nil, int(forwardJob.Status), 0).(int)
		if forwardJob != nil {
			entity.Status = int(forwardJob.Status)
		} else {
			entity.Status = 0
		}
	}

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: "success", Data: pageData}

	c.ServeJSON()

}

// @router /u/AddForward [get,post]
func (c *ForwardCtrl) AddForward() {

	entity := Models.PortForward{}
	entity.Status = 1

	c.Data["entity"] = entity

	c.TplName = "ucenter/forwardForm.html"

}

// @router /u/EditForward [get,post]
func (c *ForwardCtrl) EditForward() {

	id, _ := c.GetInt("id")

	entity := Service.SysDataS.GetPortForwardById(id)
	c.Data["entity"] = entity

	c.TplName = "ucenter/forwardForm.html"

}

// @router /u/DelForward [post]
func (c *ForwardCtrl) DelForward() {

	ids := c.GetString("ids")

	var idArray []int
	for _, id := range Utils.Split(ids, ",") {
		_id := Utils.ToInt(id)

		//检查是否正在转发中
		entity := Service.SysDataS.GetPortForwardById(_id)
		forwardJob := Service.SysDataS.GetForwardJob(entity)
		if forwardJob != nil && forwardJob.Status == Constant.RunStatus_Running {
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: fmt.Sprint("[", entity.Name, "] 正在转发中，不能删除")}
			c.ServeJSON()
			return
		} else {
			idArray = append(idArray, _id)
		}

	}

	err := Service.SysDataS.DelPortForwards(idArray)
	if err == nil {
		//
		c.Data["json"] = Models.FuncResult{Code: 0, Msg: "success"}
	} else {
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: err.Error()}

		logs.Error("DelForward err：", err)
	}

	c.ServeJSON()
}

// @router /u/SaveForward [post]
func (c *ForwardCtrl) SaveForward() {

	id, _ := c.GetInt("id")
	name := c.GetString("name", "")
	addr := c.GetString("addr", "")
	port, _ := c.GetInt("port")
	protocol := c.GetString("protocol", "TCP")
	targetAddr := c.GetString("targetAddr", "")
	targetPort, _ := c.GetInt("targetPort")
	others := c.GetString("others", "")
	fType, _ := c.GetInt("fType")
	status, _ := c.GetInt("status")

	if Utils.IsEmpty(name) {
		//
		//c.Data["json"] = Models.FuncResult{Code: 1, Msg: "名称 不能为空"}
		//c.ServeJSON()
		//return
		name = "-"
	}

	if port < 0 || port > 65535 {
		//
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "监听端口 不在允许的范围"}
		c.ServeJSON()
		return
	}

	if Utils.IsEmpty(targetAddr) {
		//
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "目标地址 不能为空"}
		c.ServeJSON()
		return
	}

	if targetPort < 0 || targetPort > 65535 {
		//
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "目标端口 不在允许的范围"}
		c.ServeJSON()
		return
	}

	if status != 0 && status != 1 {
		//
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "输入的 启用/禁用 值不正确"}
		c.ServeJSON()
		return
	}

	// if Utils.IsNotEmpty(others) {
	// 	//如果有others信息，则检查

	// }

	if fType > 0 {
		//内网穿透模式，暂不支持多端口分发
		others = ""
	}

	if id > 0 {
		entity := Service.SysDataS.GetPortForwardById(id)
		forwardJob := Service.SysDataS.GetForwardJob(entity)
		if forwardJob != nil && forwardJob.Status == Constant.RunStatus_Running {
			//正在转发中，修改前先关闭
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: fmt.Sprint("[", entity.Name, "] 正在转发中，不能修改")}
			c.ServeJSON()
			return
		}
	}

	name = Utils.FilterHtml(name)

	entity := &Models.PortForward{}
	entity.Id = id
	entity.Name = name
	entity.Addr = addr
	entity.Port = port
	entity.Protocol = protocol
	entity.TargetAddr = targetAddr
	entity.TargetPort = targetPort
	entity.Others = others
	entity.FType = fType
	entity.Status = status

	err := Service.SysDataS.SavePortForward(entity)
	if err == nil {
		c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}
	} else {
		logs.Error("SaveForward ", err.Error())
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: err.Error()}
	}

	c.ServeJSON()
}

// @router /u/OpenForward [get,post]
func (c *ForwardCtrl) OpenForward() {

	id, _ := c.GetInt("id")
	entity := Service.SysDataS.GetPortForwardById(id)

	resultChan := make(chan Models.FuncResult)
	config := Service.SysDataS.ToForwardConfig(entity)
	go Service.ForWardServ.OpenForward(config, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()

}

// @router /u/CloseForward [get,post]
func (c *ForwardCtrl) CloseForward() {
	id, _ := c.GetInt("id")

	entity := Service.SysDataS.GetPortForwardById(id)
	config := Service.SysDataS.ToForwardConfig(entity)
	Service.ForWardServ.CloseForward(config)

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()
}

// @router /u/OpenAllForward [get,post]
func (c *ForwardCtrl) OpenAllForward() {

	forwards := Service.SysDataS.GetAllPortForwardList(1)
	for _, entity := range forwards {
		resultChan := make(chan Models.FuncResult)

		forwardJob := Service.SysDataS.GetForwardJob(entity)
		if forwardJob != nil && forwardJob.Status == Constant.RunStatus_Running {
			//正在转发中
			continue
		}

		config := Service.SysDataS.ToForwardConfig(entity)
		go Service.ForWardServ.OpenForward(config, resultChan)

		fmt.Println(<-resultChan)
	}

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()
}

// @router /u/CloseAllForward [get,post]
func (c *ForwardCtrl) CloseAllForward() {
	//forwards := Service.SysDataS.GetAllPortForwardList(1)
	//for _, entity := range forwards {
	//	config := Service.SysDataS.ToForwardConfig(entity)
	//	Service.ForWardServ.CloseForward(config)
	//}

	Service.ForWardServ.CloseAllForward()
	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()
}

// @router /u/ApiDoc [get]
func (c *ForwardCtrl) ApiDoc() {

	c.TplName = "ucenter/apiDoc.html"

}

// @router /u/NetAgent [get]
func (c *ForwardCtrl) NetAgent() {

	magicAddr := beego.AppConfig.DefaultString("magic.service", ":7000")
	c.Data["magicAddr"] = magicAddr

	agentForward := Service.MagicServ.ForwardInfo
	if agentForward == nil {
		agentForward = new(Models.PortForward)
		agentForward.Addr = ""
		agentForward.Port = 3307
		agentForward.Protocol = "TCP"
		agentForward.TargetAddr = "127.0.0.1"
		agentForward.TargetPort = 3306
		agentForward.FType = 2
	}
	c.Data["agentForward"] = agentForward

	c.TplName = "ucenter/netAgent.html"
}

// @router /u/OpenMagicService [post]
func (c *ForwardCtrl) OpenMagicService() {

	addr := beego.AppConfig.DefaultString("magic.service", ":7000")

	resultChan := make(chan Models.FuncResult)
	go Service.MagicServ.StartMagicService(addr, resultChan)

	c.Data["json"] = <-resultChan
	//c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()

}

// @router /u/CloseMagicService [post]
func (c *ForwardCtrl) CloseMagicService() {

	resultChan := make(chan Models.FuncResult)
	go Service.MagicServ.StopMagicService(resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()

}

// @router /u/GetMagicStatus [post]
func (c *ForwardCtrl) GetMagicStatus() {

	magicListener := Service.MagicServ.GetMagicListener()
	if magicListener == nil {
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "未运行"}
	} else {
		c.Data["json"] = Models.FuncResult{Code: 0, Msg: "正在运行中..."}
	}

	c.ServeJSON()
}

// @router /u/GetNetAgentStatus [post]
func (c *ForwardCtrl) GetNetAgentStatus() {
	agentMap := Service.MagicServ.GetMagicClient()

	if len(agentMap) > 0 {
		count := len(agentMap)
		for k, _ := range agentMap {
			c.Data["json"] = Models.FuncResult{Code: 0, Msg: k, Data: count}
			//只取1个先
			break
		}

	} else {
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "未检测到Agent连接"}
	}

	c.ServeJSON()
}

// @router /u/ClearNetAgentStatus [post]
func (c *ForwardCtrl) ClearNetAgentStatus() {

	agentMap := Service.MagicServ.GetMagicClient()

	if len(agentMap) > 0 {
		for k, v := range agentMap {
			if v != nil {
				v.Close()
				Service.MagicServ.UnRegistryMagicClient(k)
				logs.Debug("关闭Agent：", k)
			}
		}

	}

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}
	c.ServeJSON()
}

// @router /u/StartAgentJob [post]
func (c *ForwardCtrl) StartAgentJob() {

	lAddr := c.GetString("lAddr", "")
	protocol := c.GetString("protocol", "TCP")
	targetAddr := c.GetString("targetAddr", "")
	fType, _ := c.GetInt("fType")

	portForward := new(Models.PortForward)
	portForward.Addr = Utils.Split(lAddr, ":")[0]
	portForward.Port = Utils.ToInt(Utils.Split(lAddr, ":")[1])
	portForward.Protocol = protocol
	portForward.TargetAddr = Utils.Split(targetAddr, ":")[0]
	portForward.TargetPort = Utils.ToInt(Utils.Split(targetAddr, ":")[1])
	portForward.FType = fType

	resultChan := make(chan Models.FuncResult)
	go Service.MagicServ.StartMagicForward(portForward, resultChan)

	c.Data["json"] = <-resultChan
	c.ServeJSON()

}

// @router /u/StopAgentJob [post]
func (c *ForwardCtrl) StopAgentJob() {

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}
	c.ServeJSON()

}

// @router /u/ChangeForwardDebug [get,post]
func (c *ForwardCtrl) ChangeForwardDebug() {

	id, _ := c.GetInt("status")

	Service.ForWardDebug = Utils.If(id == 1, true, false).(bool)

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()

}

// @router /u/AddBatchForward [get]
func (c *ForwardCtrl) AddBatchForward() {

	c.TplName = "ucenter/addBatchForward.html"

}

// @router /u/SaveBatchForward [post]
func (c *ForwardCtrl) SaveBatchForward() {
	rows, _ := c.GetInt("rows")

	var entities []*Models.PortForward
	for i := 0; i < rows; i++ {
		name := c.GetString(fmt.Sprint("name[", i, "]"), "-")
		port, _ := c.GetInt(fmt.Sprint("port[", i, "]"))
		protocol := c.GetString(fmt.Sprint("protocol[", i, "]"), "TCP")
		targetAddr := c.GetString(fmt.Sprint("targetAddr[", i, "]"), "")
		targetPort, _ := c.GetInt(fmt.Sprint("targetPort[", i, "]"))

		if port < 0 || port > 65535 {
			//
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: fmt.Sprint("监听端口 不在允许的范围 ", port)}
			c.ServeJSON()
			return
		}

		if targetPort < 0 || targetPort > 65535 {
			//
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: fmt.Sprint("目标端口 不在允许的范围 ", targetPort)}
			c.ServeJSON()
			return
		}

		if Utils.IsEmpty(targetAddr) {
			//
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: "目标地址 不能为空"}
			c.ServeJSON()
			return
		}

		name = Utils.FilterHtml(name)

		entity := &Models.PortForward{}
		entity.Id = 0
		entity.Name = name
		entity.Addr = ""
		entity.Port = port
		entity.Protocol = protocol
		entity.TargetAddr = targetAddr
		entity.TargetPort = targetPort
		entity.Others = ""
		entity.FType = 0
		entity.Status = 1

		entities = append(entities, entity)

	}

	for _, entity := range entities {
		err := Service.SysDataS.SavePortForward(entity)
		if err != nil {
			logs.Error("SaveForward ", err.Error())
		}
	}

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()

}

// @router /u/ImportForward [get]
func (c *ForwardCtrl) ImportForward() {

	c.TplName = "ucenter/importForward.html"

}

// @router /u/SaveImportForward [post]
func (c *ForwardCtrl) SaveImportForward() {

	splitChar := c.GetString("splitChar", ",")
	inputDatas := c.GetString("inputDatas", "")

	if Utils.IsEmpty(inputDatas) {
		//
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "导入的数据不能为空"}
		c.ServeJSON()
		return
	}

	dataRows := Utils.Split(inputDatas, "\n")

	var entities []*Models.PortForward
	for _, rowContent := range dataRows {
		if Utils.IsEmpty(rowContent) {
			continue
		}

		//名称,本地监听地址,本地监听端口,协议类型,目标地址,目标端口
		rowDatas := Utils.Split(rowContent, splitChar)

		if len(rowDatas) < 6 {
			continue
		}

		name := rowDatas[0]
		name = Utils.FilterHtml(name)
		if Utils.IsEmpty(name) {
			name = "-"
		}

		addr := rowDatas[1]
		port := Utils.ToInt(rowDatas[2])
		protocol := Utils.If(rowDatas[3] == "TCP", "TCP", "UDP").(string)
		targetAddr := rowDatas[4]
		targetPort := Utils.ToInt(rowDatas[5])

		if Utils.IsEmpty(targetAddr) {
			//
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: "目标地址 不能为空"}
			c.ServeJSON()
			return
		}

		if port < 0 || port > 65535 {
			//
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: fmt.Sprint("监听端口 不在允许的范围 ", port)}
			c.ServeJSON()
			return
		}

		if targetPort < 0 || targetPort > 65535 {
			//
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: fmt.Sprint("目标端口 不在允许的范围 ", targetPort)}
			c.ServeJSON()
			return
		}

		entity := &Models.PortForward{}
		entity.Id = 0
		entity.Name = name
		entity.Addr = addr
		entity.Port = port
		entity.Protocol = protocol
		entity.TargetAddr = targetAddr
		entity.TargetPort = targetPort
		entity.Others = ""
		entity.FType = 0
		entity.Status = 1

		entities = append(entities, entity)

	}

	for _, entity := range entities {
		err := Service.SysDataS.SavePortForward(entity)
		if err != nil {
			logs.Error("SaveForward ", err.Error())
		}
	}

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()
}
