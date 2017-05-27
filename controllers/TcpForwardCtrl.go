package controllers

import (
	"port-forward/controllers/base"
	"port-forward/models"
	"port-forward/services"
)

type TcpForwardCtrl struct {
	BaseController.ConsoleController
}

// @router /u/TcpForwardList [get]
func (c *TcpForwardCtrl) TcpForwardList() {
	c.TplName = "ucenter/index.html"
}

// @router /u/OpenTcpForward [get,post]
func (c *TcpForwardCtrl) OpenTcpForward() {
	fromAddr := c.GetString("fromAddr")
	toAddr := c.GetString("toAddr")

	resultChan := make(chan models.ResultData)
	go services.ForwardS.StartTcpPortForward(fromAddr, toAddr, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()

}

// @router /u/CloseTcpForward [get,post]
func (c *TcpForwardCtrl) CloseTcpForward() {
	fromAddr := c.GetString("fromAddr")

	resultChan := make(chan models.ResultData)
	go services.ForwardS.CloseTcpPortForward(fromAddr, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()
}
