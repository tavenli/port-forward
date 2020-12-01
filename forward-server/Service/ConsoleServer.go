package Service

import (
	"encoding/gob"
	"forward-core/Models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ConsoleServer struct {

}

func (_self *ConsoleServer) StartHttpServer() {

	//开启seesion支持，默认使用的存储引擎为：memory
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	//默认static目录是可以直接访问的，其它目录需要单独指定
	beego.SetStaticPath("/theme", "theme")

	//
	gob.Register(&Models.LoginUser{})

	logs.Debug("Http 服务启动...")

	//启动应用
	beego.Run()


}
