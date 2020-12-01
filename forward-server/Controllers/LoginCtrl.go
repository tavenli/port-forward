package Controllers

import (
	"forward-core/Models"
	"forward-core/NetUtils"
	"forward-core/Utils"
	"forward-server/Controllers/BaseCtrl"
	"forward-server/Service"
	"github.com/astaxie/beego/logs"
)

type LoginCtrl struct {
	BaseCtrl.ConsoleCtrl
}

// @router /logout
func (c *LoginCtrl) Logout() {
	c.ClearUserInfo()
	c.Ctx.Redirect(302, "/login")

}

// @router /login [get]
func (c *LoginCtrl) Login() {

	c.TplName = "login.html"

}

// @router /login [post]
func (c *LoginCtrl) DoLogin() {
	userName := c.GetString("userName")
	passWord := c.GetString("passWord")

	sysUser := Service.SysDataS.GetSysUserByName(userName)
	if sysUser == nil {
		logs.Debug("用户不存在")
		c.Ctx.Redirect(302, "/login")
		return
	}

	descryptPwd := Utils.GetMd5(passWord)
	logs.Debug("存储的密码：", sysUser.PassWord, " 输入的密码：", descryptPwd)
	if sysUser.PassWord == descryptPwd {
		logs.Info("用户登录：", userName, " IP：", NetUtils.GetIP(&c.Controller))
		loginUser := new(Models.LoginUser)
		loginUser.UserId = 1
		loginUser.UserName = userName

		c.SetSession("userInfo", loginUser)
		c.Ctx.Redirect(302, "/u/main")
	} else {
		logs.Debug("用户登录失败")
		c.Ctx.Redirect(302, "/login")
	}

}
