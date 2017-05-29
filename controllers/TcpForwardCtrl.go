package controllers

import (
	"port-forward/controllers/base"
	"port-forward/models"
	"port-forward/services"
	"port-forward/utils"

	"fmt"

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

	if utils.IsEmpty(name) {
		//
		c.Data["json"] = models.ResultData{Code: 1, Msg: "名称 不能为空"}
		c.ServeJSON()
		return
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

	fromAddr := fmt.Sprint(entity.Addr, ":", entity.Port)
	toAddr := fmt.Sprint(entity.TargetAddr, ":", entity.TargetPort)

	resultChan := make(chan models.ResultData)
	go services.ForwardS.StartPortForward(fromAddr, toAddr, resultChan)

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
	go services.ForwardS.ClosePortForward(fromAddr, toAddr, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()
}
