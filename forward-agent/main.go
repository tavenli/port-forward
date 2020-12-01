package main

import (
	"forward-core/Utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main()  {

	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	//为了让日志输出不影响性能，开启异步日志
	logs.Async()

	//远端外网服务
	magicServerAddr := beego.AppConfig.String("magic.server")

	logs.Debug("★★★★★★★★★★★★★★★★★★★★")
	logs.Debug("         tcp-forward 启动")
	logs.Debug("")
	logs.Debug("项目地址：https://github.com/tavenli/port-forward")
	logs.Debug("")
	logs.Debug("请求远端控制服务器：", magicServerAddr)
	logs.Debug("")
	logs.Debug("★★★★★★★★★★★★★★★★★★★★")

	if Utils.IsEmpty(magicServerAddr){
		magicServerAddr = "forward.apiclub.top:7000"
	}

	agentService := new(AgentServiceV1)
	agentService.MagicServerAddr = magicServerAddr
	agentService.AgentOnline = false

	go agentService.ConnToMagicServer()

	go agentService.ConnectToMagicLoop()

	select {

	}


}