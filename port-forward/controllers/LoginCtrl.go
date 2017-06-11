package controllers

import (
	"port-forward/controllers/base"
	"port-forward/models"
	"port-forward/utils"

	"port-forward/services"

	"github.com/astaxie/beego/logs"
)

type LoginCtrl struct {
	BaseController.ConsoleController
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

	sysUser := services.SysDataS.GetSysUserByName(userName)
	if sysUser == nil {
		logs.Debug("用户不存在")
		c.Ctx.Redirect(302, "/login")
		return
	}
	descryptPwd := utils.GetMd5(passWord)
	logs.Debug("存储的密码：", sysUser.PassWord, " 输入的密码：", descryptPwd)
	if sysUser.PassWord == descryptPwd {
		logs.Info("用户登录：", userName, " IP：", utils.GetIP(&c.Controller))
		loginUser := new(models.LoginUser)
		loginUser.UserId = 1
		loginUser.UserName = userName

		c.SetSession("userInfo", loginUser)
		c.Ctx.Redirect(302, "/u/main")
	} else {
		logs.Debug("用户登录失败")
		c.Ctx.Redirect(302, "/login")
	}
}
