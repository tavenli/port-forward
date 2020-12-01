package BaseCtrl

import (
	"forward-core/NetUtils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ApiCtrl struct {
	beego.Controller
}

var apiAuth = beego.AppConfig.String("api.auth")

func (c *ApiCtrl) Prepare() {
	reqUrl := c.Ctx.Request.RequestURI
	userIp := NetUtils.GetIP(&c.Controller)

	logs.Debug("执行Prepare，当前reqUrl：", reqUrl, " userIp:", userIp)

	//校验鉴权参数
	auth := c.GetString("auth")

	if auth != apiAuth {
		//logs.Error("apiAuth验证失败：", auth)
		//c.Ctx.Redirect(302, "/apiAuthFail")
	}

}