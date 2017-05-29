package controllers

import (
	"port-forward/controllers/base"
	"port-forward/utils"
	"runtime"
	"time"
)

type UCenterCtrl struct {
	BaseController.ConsoleController
}

// @router /u/main [get]
func (c *UCenterCtrl) Main() {

	c.Layout = "ucenter/layout.html"
	c.TplName = "ucenter/main.html"

}

// @router /u/index [get]
func (c *UCenterCtrl) Index() {

	c.Data["runtime_NumCPU"] = runtime.NumCPU()
	c.Data["runtime_GOOS"] = runtime.GOOS
	c.Data["runtime_GOARCH"] = runtime.GOARCH
	c.Data["runtime_NumGoroutine"] = runtime.NumGoroutine()
	c.Data["server_Time"] = time.Now()

	c.TplName = "ucenter/index.html"
}

// @router /u/getServerTime [post]
func (c *UCenterCtrl) GetServerTime() {

	c.Data["json"] = utils.GetCurrentTime()

	c.ServeJSON()

}
