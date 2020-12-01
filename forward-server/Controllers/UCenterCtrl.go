package Controllers

import (
	"forward-core/Models"
	"forward-core/Utils"
	"forward-server/Controllers/BaseCtrl"
	"forward-server/Service"
	"runtime"
	"time"
)

type UCenterCtrl struct {
	BaseCtrl.ConsoleCtrl
}

// @router /u/main [get]
func (c *UCenterCtrl) Main() {

	c.Layout = "ucenter/layout.html"
	c.TplName = "ucenter/main.html"

}

// @router /u/index [get]
func (c *UCenterCtrl) Index() {

	c.Data["runtime_NumCPU"] = runtime.NumCPU()
	c.Data["runtime_GOOS"] = runtime.GOOS
	c.Data["runtime_GOARCH"] = runtime.GOARCH
	c.Data["runtime_NumGoroutine"] = runtime.NumGoroutine()
	c.Data["server_Time"] = time.Now()

	c.TplName = "ucenter/index.html"
}

// @router /u/getServerTime [post]
func (c *UCenterCtrl) GetServerTime(){

	c.Data["json"] = Utils.GetCurrentTime()

	c.ServeJSON()

}

// @router /u/changePwd [get]
func (c *UCenterCtrl) ChangePwd() {

	c.TplName = "ucenter/changePwd.html"
}

// @router /u/doChangePwd [post]
func (c *UCenterCtrl) DoChangePwd() {
	userInfo := c.GetUserInfo()

	passWord := c.GetString("passWord")
	passWord2 := c.GetString("passWord2")

	if Utils.IsEmpty(passWord) {
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "密码不能为空"}
		c.ServeJSON()
		return
	}

	if passWord != passWord2 {
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: "两次输入的密码不一致"}
		c.ServeJSON()
		return
	}

	err := Service.SysDataS.ChangeUserPwd(userInfo.UserId, passWord)
	if err == nil {
		c.Data["json"] = Models.FuncResult{Code: 0, Msg: "密码修改成功"}
	} else {
		c.Data["json"] = Models.FuncResult{Code: 1, Msg: err.Error()}
	}

	c.ServeJSON()

}
