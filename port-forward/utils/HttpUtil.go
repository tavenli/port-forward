package utils

import "github.com/astaxie/beego"

func GetIP(c *beego.Controller) string {
	//也可以直接用 c.Ctx.Input.IP() 取真实IP
	ip := c.Ctx.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip = c.Ctx.Request.Header.Get("Remote_addr")
	if ip == "" {
		ip = c.Ctx.Request.RemoteAddr
	}
	return ip
}
