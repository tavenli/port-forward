package BaseController

import (
	"port-forward/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ConsoleController struct {
	beego.Controller
}

var (
	ConsoleLoginUrl string = "/login"
)

func (c *ConsoleController) Prepare() {
	reqUrl := c.Ctx.Request.RequestURI
	logs.Debug("执行Prepare，当前reqUrl：", reqUrl)

	if ConsoleLoginUrl == reqUrl {
		//如果是登录地址，则不校验
		return
	}

	//开始访问每个action前，执行登录和权限检查
	userInfo := c.GetUserInfo()
	if userInfo == nil {
		//未登录
		c.Ctx.Redirect(302, ConsoleLoginUrl)
	}

}

func (c *ConsoleController) StoreUserInfo(loginUser *models.LoginUser) {
	c.SetSession("userInfo", loginUser)
}

func (c *ConsoleController) GetUserInfo() *models.LoginUser {

	userInfo := c.GetSession("userInfo")
	if userInfo == nil {
		return nil
	}
	return userInfo.(*models.LoginUser)

}

func (c *ConsoleController) ClearUserInfo() {

	c.DelSession("userInfo")

}
