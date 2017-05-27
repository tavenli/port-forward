package controllers

import (
	"port-forward/controllers/base"
	"port-forward/models"
	"port-forward/services"
	"runtime"
	"time"
)

type RestApiCtrl struct {
	BaseController.ApiController
}

// @router /ServerSummary [get,post]
func (c *RestApiCtrl) ServerSummary() {
	obj := make(map[string]interface{})
	obj["runtime_NumGoroutine"] = runtime.NumGoroutine()
	obj["runtime_GOOS"] = runtime.GOOS
	obj["runtime_GOARCH"] = runtime.GOARCH
	obj["server_Time"] = time.Now()
	c.Data["json"] = obj

	c.ServeJSON()

}

// @router /OpenTcpForward [get,post]
func (c *RestApiCtrl) OpenTcpForward() {
	fromAddr := c.GetString("fromAddr")
	toAddr := c.GetString("toAddr")

	//测试
	//http://127.0.0.1:8000/api/v1/OpenTcpForward?auth=26CCD056107481F45D1AC805A24A9E59&fromAddr=:8010&toAddr=127.0.0.1:3306
	resultChan := make(chan models.ResultData)
	go services.ForwardS.StartTcpPortForward(fromAddr, toAddr, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()

}

// @router /CloseTcpForward [get,post]
func (c *RestApiCtrl) CloseTcpForward() {
	fromAddr := c.GetString("fromAddr")

	//测试
	//http://127.0.0.1:8000/api/v1/CloseTcpForward?auth=26CCD056107481F45D1AC805A24A9E59&fromAddr=:8010
	resultChan := make(chan models.ResultData)
	go services.ForwardS.CloseTcpPortForward(fromAddr, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()
}
