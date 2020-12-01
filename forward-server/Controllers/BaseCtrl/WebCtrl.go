package BaseCtrl

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type WebCtrl struct {
	beego.Controller
}

func (c *WebCtrl) Prepare() {
	reqUrl := c.Ctx.Request.RequestURI
	logs.Debug("执行Prepare，当前reqUrl：", reqUrl)

}
