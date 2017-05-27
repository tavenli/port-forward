package BaseController

import (
	"port-forward/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ApiController struct {
	beego.Controller
}

var apiAuth = beego.AppConfig.String("api.auth")

func (c *ApiController) Prepare() {
	reqUrl := c.Ctx.Request.RequestURI
	userIp := utils.GetIP(&c.Controller)
	logs.Debug("执行Prepare，当前reqUrl：", reqUrl, " userIp:", userIp)

	//校验鉴权参数
	auth := c.GetString("auth")

	if auth != apiAuth {
		logs.Error("apiAuth验证失败：", auth)
		c.Ctx.Redirect(302, "/apiAuthFail")
	}

}
