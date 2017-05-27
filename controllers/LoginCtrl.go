package controllers

import (
	"port-forward/controllers/base"
	"port-forward/models"
)

type LoginCtrl struct {
	BaseController.WebController
}

// @router /logout
func (c *LoginCtrl) Logout() {

	c.Ctx.Redirect(302, "/login")

}

// @router /login [get]
func (c *LoginCtrl) Login() {

	c.TplName = "login.html"

}

// @router /login [post]
func (c *LoginCtrl) DoLogin() {

}

// @router /apiAuthFail [get]
func (c *LoginCtrl) ApiAuthFail() {

	c.Data["json"] = models.ResultData{Code: 1, Msg: "ApiAuth鉴权失败"}

	c.ServeJSON()

}
