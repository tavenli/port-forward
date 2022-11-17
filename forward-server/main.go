package main

import (
	"fmt"
	"forward-server/Service"
	_ "forward-server/routers"

	"forward-core/Models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vmihailenco/msgpack"
)

func main() {

	logFileConfig := beego.AppConfig.String("logfile.config")

	//日志级别："emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"
	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)

	if len(logFileConfig) == 0 {
		//logFileConfig = `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`
		logFileConfig = `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"separate":["error"]}`

	}

	if logFileConfig != "close" {
		//logs.SetLogger(logs.AdapterFile, `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
		//logs.SetLogger(logs.AdapterFile, logFileConfig)
		logs.SetLogger(logs.AdapterMultiFile, logFileConfig)
	}

	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	//为了让日志输出不影响性能，开启异步日志
	logs.Async()

	logs.Debug("★★★★★★★★★★★★★★★★★★★★")
	logs.Debug("         port-forward 启动")
	logs.Debug("")
	logs.Debug("项目地址：https://github.com/tavenli/port-forward")
	logs.Debug("")
	logs.Debug("★★★★★★★★★★★★★★★★★★★★")

	defer logs.GetBeeLogger().Flush()

	//test1()

	//启动Web控制台和接口
	Service.ConsoleServ.StartHttpServer()

	//select {}

	//endRunning := make(chan bool, 1)
	//time.Sleep(1* time.Second)
	//endRunning <- true
	//<-endRunning
}

func test1() {

	//github.com/gogf/gf/g/os/glog
	//glog.Debug("This is Debug")
	//glog.Info("This is Info")

	b, err := msgpack.Marshal(&Models.PortInfo{Addr: "bar"})
	if err != nil {
		panic(err)
	}

	var item Models.PortInfo
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}

	logs.Debug(item.Addr)

	//
	config := new(Models.ForwardConfig)
	//config.Protocol = "TCP"
	config.Protocol = "UDP"
	config.SrcAddr = ""
	config.SrcPort = 8888
	//106.14.184.192:9999
	//config.DestAddr = "106.14.184.192"
	//config.DestPort = 9999
	config.DestAddr = "svn.apiclub.top"
	config.DestPort = 9900
	config.Status = 0
	config.Name = "测试1"
	config.RuleId = 1

	resultChan := make(chan Models.FuncResult)
	Service.ForWardServ.OpenForward(config, resultChan)

	fmt.Println(<-resultChan)

}
