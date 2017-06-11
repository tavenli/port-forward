package controllers

import (
	"port-forward/controllers/base"
	"port-forward/models"
	"port-forward/services"
	"port-forward/utils"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ForwardCtrl struct {
	BaseController.ConsoleController
}

// @router /u/ForwardList [get]
func (c *ForwardCtrl) ForwardList() {

	c.TplName = "ucenter/forwardList.html"
}

// @router /u/ForwardList/json [post]
func (c *ForwardCtrl) ForwardListJson() {
	pageParam := new(models.PageParam)
	pageParam.PIndex, _ = c.GetInt64("pIndex")
	pageParam.PSize, _ = c.GetInt64("pSize")

	port, _ := c.GetInt("port")

	query := &models.PortForward{}
	query.Port = port
	query.FType = -1

	pageData := services.SysDataS.GetPortForwardList(query, pageParam.PIndex, pageParam.PSize)

	for _, entity := range pageData.Data.([]*models.PortForward) {
		key := services.ForwardS.GetKeyByEntity(entity)
		entity.Status = utils.If(services.ForwardS.PortConflict(key), 1, 0).(int)
	}

	c.Data["json"] = models.ResultData{Code: 0, Msg: "success", Data: pageData}

	c.ServeJSON()

}

// @router /u/AddForward [get,post]
func (c *ForwardCtrl) AddForward() {

	entity := models.PortForward{}

	c.Data["entity"] = entity

	c.TplName = "ucenter/forwardForm.html"

}

// @router /u/EditForward [get,post]
func (c *ForwardCtrl) EditForward() {

	id, _ := c.GetInt("id")

	entity := services.SysDataS.GetPortForwardById(id)
	c.Data["entity"] = entity

	c.TplName = "ucenter/forwardForm.html"

}

// @router /u/DelForward [post]
func (c *ForwardCtrl) DelForward() {

	ids := c.GetString("ids")

	var idArray []int
	for _, id := range utils.Split(ids, ",") {
		_id := utils.ToInt(id)

		//检查是否正在转发中
		entity := services.SysDataS.GetPortForwardById(_id)
		key := services.ForwardS.GetKeyByEntity(entity)
		if services.ForwardS.PortConflict(key) {
			c.Data["json"] = models.ResultData{Code: 1, Msg: fmt.Sprint("[", entity.Name, "] 正在转发中，不能删除")}
			c.ServeJSON()
			return
		} else {
			idArray = append(idArray, _id)
		}

	}

	err := services.SysDataS.DelPortForwards(idArray)
	if err == nil {
		//
		c.Data["json"] = models.ResultData{Code: 0, Msg: "success"}
	} else {
		c.Data["json"] = models.ResultData{Code: 1, Msg: err.Error()}

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
	//protocol := c.GetString("protocol", "TCP")
	targetAddr := c.GetString("targetAddr", "")
	targetPort, _ := c.GetInt("targetPort")
	others := c.GetString("others", "")
	fType, _ := c.GetInt("fType")

	if utils.IsEmpty(name) {
		//
		c.Data["json"] = models.ResultData{Code: 1, Msg: "名称 不能为空"}
		c.ServeJSON()
		return
	}

	if port < 0 || port > 65535 {
		//
		c.Data["json"] = models.ResultData{Code: 1, Msg: "监听端口 不在允许的范围"}
		c.ServeJSON()
		return
	}

	if utils.IsEmpty(targetAddr) {
		//
		c.Data["json"] = models.ResultData{Code: 1, Msg: "目标地址 不能为空"}
		c.ServeJSON()
		return
	}

	if targetPort < 0 || targetPort > 65535 {
		//
		c.Data["json"] = models.ResultData{Code: 1, Msg: "目标端口 不在允许的范围"}
		c.ServeJSON()
		return
	}

	// if utils.IsNotEmpty(others) {
	// 	//如果有others信息，则检查

	// }

	if fType > 0 {
		//内网穿透模式，暂不支持多端口分发
		others = ""
	}

	if id > 0 {
		entity := services.SysDataS.GetPortForwardById(id)
		key := services.ForwardS.GetKeyByEntity(entity)
		if services.ForwardS.PortConflict(key) {
			//正在转发中，修改前先关闭
			fromAddr := fmt.Sprint(entity.Addr, ":", entity.Port)
			toAddr := fmt.Sprint(entity.TargetAddr, ":", entity.TargetPort)
			resultChan := make(chan models.ResultData)
			go services.ForwardS.ClosePortForward(fromAddr, toAddr, entity.FType, resultChan)
		}
	}

	name = utils.FilterHtml(name)

	entity := &models.PortForward{}
	entity.Id = id
	entity.Name = name
	entity.Addr = addr
	entity.Port = port
	entity.Protocol = "TCP"
	entity.TargetAddr = targetAddr
	entity.TargetPort = targetPort
	entity.Others = others
	entity.FType = fType

	err := services.SysDataS.SavePortForward(entity)
	if err == nil {
		c.Data["json"] = models.ResultData{Code: 0, Msg: ""}
	} else {
		logs.Error("SaveForward ", err.Error())
		c.Data["json"] = models.ResultData{Code: 1, Msg: err.Error()}
	}

	c.ServeJSON()
}

// @router /u/OpenForward [get,post]
func (c *ForwardCtrl) OpenForward() {
	id, _ := c.GetInt("id")
	entity := services.SysDataS.GetPortForwardById(id)

	resultChan := make(chan models.ResultData)
	go services.ForwardS.StartPortForward(entity, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()

}

// @router /u/CloseForward [get,post]
func (c *ForwardCtrl) CloseForward() {
	id, _ := c.GetInt("id")

	entity := services.SysDataS.GetPortForwardById(id)

	fromAddr := fmt.Sprint(entity.Addr, ":", entity.Port)
	toAddr := fmt.Sprint(entity.TargetAddr, ":", entity.TargetPort)

	resultChan := make(chan models.ResultData)
	go services.ForwardS.ClosePortForward(fromAddr, toAddr, entity.FType, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()
}

// @router /u/ApiDoc [get]
func (c *ForwardCtrl) ApiDoc() {

	c.TplName = "ucenter/apiDoc.html"

}

// @router /u/NetAgent [get]
func (c *ForwardCtrl) NetAgent() {

	c.TplName = "ucenter/netAgent.html"
}

// @router /u/OpenMagicService [post]
func (c *ForwardCtrl) OpenMagicService() {

	addr := beego.AppConfig.DefaultString("magic.service", ":7000")

	resultChan := make(chan models.ResultData)
	go services.ForwardS.StartMagicService(addr, resultChan)

	c.Data["json"] = <-resultChan
	//c.Data["json"] = models.ResultData{Code: 0, Msg: ""}

	c.ServeJSON()
}

// @router /u/CloseMagicService [post]
func (c *ForwardCtrl) CloseMagicService() {

	resultChan := make(chan models.ResultData)
	go services.ForwardS.StopMagicService(resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()
}

// @router /u/GetMagicStatus [post]
func (c *ForwardCtrl) GetMagicStatus() {

	magicListener := services.ForwardS.GetMagicListener()
	if magicListener == nil {
		c.Data["json"] = models.ResultData{Code: 1, Msg: "未运行"}
	} else {
		c.Data["json"] = models.ResultData{Code: 0, Msg: "正在运行中..."}
	}

	c.ServeJSON()
}

// @router /u/GetNetAgentStatus [post]
func (c *ForwardCtrl) GetNetAgentStatus() {
	agentMap := services.ForwardS.GetMagicClient()

	if len(agentMap) > 0 {
		count := len(agentMap)
		for k, _ := range agentMap {
			c.Data["json"] = models.ResultData{Code: 0, Msg: k, Data: count}
			//只取1个先
			break
		}

	} else {
		c.Data["json"] = models.ResultData{Code: 1, Msg: "未检测到Agent连接"}
	}

	c.ServeJSON()
}

// @router /u/ClearNetAgentStatus [post]
func (c *ForwardCtrl) ClearNetAgentStatus() {
	agentMap := services.ForwardS.GetMagicClient()

	if len(agentMap) > 0 {
		for k, v := range agentMap {
			if v != nil {
				v.Close()
				services.ForwardS.UnRegistryMagicClient(k)
				logs.Debug("关闭Agent：", k)
			}
		}

	}

	c.Data["json"] = models.ResultData{Code: 0, Msg: ""}
	c.ServeJSON()
}
