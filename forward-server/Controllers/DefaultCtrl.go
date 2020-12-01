package Controllers

import (
	"forward-core/Models"
	"forward-server/Controllers/BaseCtrl"
)

type DefaultCtrl struct {
	BaseCtrl.WebCtrl
}


// @router / [get]
func (c *DefaultCtrl) Default() {

	c.Ctx.Redirect(302, "/login")

	//c.Data["currentTime"] = time.Now()
	//c.TplName = "index.html"
}

// @router /apiAuthFail [get]
func (c *DefaultCtrl) ApiAuthFail() {

	c.Data["json"] = Models.FuncResult{Code: 1, Msg: "ApiAuth鉴权失败"}

	c.ServeJSON()

}