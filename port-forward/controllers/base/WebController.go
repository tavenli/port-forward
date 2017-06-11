package BaseController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) Prepare() {
	reqUrl := c.Ctx.Request.RequestURI
	logs.Debug("执行Prepare，当前reqUrl：", reqUrl)

}
