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

// @router /OpenForward [get,post]
func (c *RestApiCtrl) OpenForward() {
	fromAddr := c.GetString("fromAddr")
	toAddr := c.GetString("toAddr")

	entity := services.SysDataS.ChkPortForwardByApi(fromAddr, "TCP", toAddr)
	if entity == nil {
		_, err := services.SysDataS.SavePortForwardByApi(fromAddr, "TCP", toAddr)
		if err != nil {
			c.Data["json"] = models.ResultData{Code: 1, Msg: "保存端口配置失败"}
			c.ServeJSON()
			return
		}
	}
	//测试
	//http://127.0.0.1:8000/api/v1/OpenForward?auth=26CCD056107481F45D1AC805A24A9E59&fromAddr=:8010&toAddr=127.0.0.1:3306
	resultChan := make(chan models.ResultData)
	go services.ForwardS.StartPortForward(fromAddr, toAddr, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()

}

// @router /CloseForward [get,post]
func (c *RestApiCtrl) CloseForward() {
	fromAddr := c.GetString("fromAddr")
	toAddr := c.GetString("toAddr")

	//测试
	//http://127.0.0.1:8000/api/v1/CloseForward?auth=26CCD056107481F45D1AC805A24A9E59&fromAddr=:8010&toAddr=127.0.0.1:3306
	resultChan := make(chan models.ResultData)
	go services.ForwardS.ClosePortForward(fromAddr, toAddr, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()
}
