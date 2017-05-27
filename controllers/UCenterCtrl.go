package controllers

import "port-forward/controllers/base"

type UCenterCtrl struct {
	BaseController.ConsoleController
}

// @router /u/main [get]
func (c *UCenterCtrl) Main() {

	c.Layout = "ucenter/layout.html"
	c.TplName = "ucenter/main.html"

}

// @router /u/index [get]
func (c *UCenterCtrl) Index() {
	c.TplName = "ucenter/index.html"
}
