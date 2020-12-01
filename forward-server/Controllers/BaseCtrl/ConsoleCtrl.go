package BaseCtrl

import (
	"forward-core/Models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	ConsoleLoginUrl string = "/login"
)

type ConsoleCtrl struct {
	beego.Controller
	LoginUser *Models.LoginUser
}

func (c *ConsoleCtrl) Prepare() {
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

	c.LoginUser = userInfo

}

//判断用户是否登录.
func (c *ConsoleCtrl) isUserLoggedIn() bool {
	return c.LoginUser != nil && c.LoginUser.UserId > 0
}

func (c *ConsoleCtrl) StoreUserInfo(loginUser *Models.LoginUser) {
	c.SetSession("userInfo", loginUser)
}

func (c *ConsoleCtrl) GetUserInfo() *Models.LoginUser {

	userInfo := c.GetSession("userInfo")
	if userInfo == nil {
		return nil
	}
	return userInfo.(*Models.LoginUser)

}

func (c *ConsoleCtrl) ClearUserInfo() {

	c.DelSession("userInfo")
	c.LoginUser = nil

}