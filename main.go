package main

import (
	"encoding/gob"
	"port-forward/models"
	_ "port-forward/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	//为了让日志输出不影响性能，开启异步日志
	logs.Async()

	//开启seesion支持，默认使用的存储引擎为：memory
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	//beego.BConfig.WebConfig.Session.SessionProvider = "file"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "./session"
	gob.Register(&models.LoginUser{})

	logs.Debug("★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★")
	logs.Debug("               tcp-forward 启动")
	logs.Debug("")
	logs.Debug("开源项目地址：https://github.com/tavenli/port-forward")
	logs.Debug("")
	logs.Debug("★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★")

	//默认static目录是可以直接访问的，其它目录需要单独指定
	beego.SetStaticPath("/theme", "theme")

	//启动应用
	beego.Run()

}
