package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultCtrl struct {
	beego.Controller
}

func (c *DefaultCtrl) Get() {
	c.Data["Website"] = "App"
	c.TplName = "index.html"
}
