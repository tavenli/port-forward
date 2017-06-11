package controllers

import (
	"net"
	"net/http"
	"port-forward/controllers/base"
	"time"

	"github.com/astaxie/beego/logs"
)

type HelpCtrl struct {
	BaseController.ApiController
}

// @router /GetTcp [get,post]
func (c *HelpCtrl) GetTcp() {
	//var w http.ResponseWriter
	//c.Ctx.ResponseWriter.Hijack()
	w := c.Ctx.ResponseWriter.ResponseWriter
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		logs.Error("获取Hijacks失败:", err)
		return
	}
	if tcp, ok := conn.(*net.TCPConn); ok {
		tcp.SetKeepAlivePeriod(60 * time.Minute)
	}

	conn.Close()
}
