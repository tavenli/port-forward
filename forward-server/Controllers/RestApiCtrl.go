package Controllers

import (
	"forward-core/Models"
	"forward-core/Utils"
	"forward-server/Controllers/BaseCtrl"
	"forward-server/Service"
	"runtime"
	"strings"
	"time"
)

type RestApiCtrl struct {
	BaseCtrl.ApiCtrl
}

// @router /ServerSummary [get,post]
func (c *RestApiCtrl) ServerSummary() {
	obj := make(map[string]interface{})
	obj["runtime_NumGoroutine"] = runtime.NumGoroutine()
	obj["runtime_GOOS"] = runtime.GOOS
	obj["runtime_GOARCH"] = runtime.GOARCH
	obj["server_Time"] = time.Now()

	obj["forwardList"] = Service.ForWardServ.FindAllForward()

	c.Data["json"] = obj

	c.ServeJSON()

}

// @router /OpenForward [get,post]
func (c *RestApiCtrl) OpenForward() {

	fromAddr := c.GetString("fromAddr")
	toAddr := c.GetString("toAddr")
	protocol := c.GetString("protocol", "TCP")

	entity := Service.SysDataS.ChkPortForwardByApi(fromAddr, protocol, toAddr)
	if entity == nil {
		var err error
		entity, err = Service.SysDataS.SavePortForwardByApi(fromAddr, protocol, toAddr)
		if err != nil {
			c.Data["json"] = Models.FuncResult{Code: 1, Msg: "保存端口配置失败"}
			c.ServeJSON()
			return
		}
	}
	//测试
	//http://127.0.0.1:8000/api/v1/OpenForward?auth=26CCD056107481F45D1AC805A24A9E59&fromAddr=:8010&toAddr=127.0.0.1:3306
	resultChan := make(chan Models.FuncResult)
	config := Service.SysDataS.ToForwardConfig(entity)
	go Service.ForWardServ.OpenForward(config, resultChan)

	c.Data["json"] = <-resultChan

	c.ServeJSON()

}

// @router /CloseForward [get,post]
func (c *RestApiCtrl) CloseForward() {

	fromAddr := c.GetString("fromAddr")
	toAddr := c.GetString("toAddr")
	protocol := c.GetString("protocol", "TCP")
	//fType, _ := c.GetInt("fType", 0)

	//测试
	//http://127.0.0.1:8000/api/v1/CloseForward?auth=26CCD056107481F45D1AC805A24A9E59&fromAddr=:8010&toAddr=127.0.0.1:3306

	config := new(Models.ForwardConfig)
	config.RuleId = 0
	config.Name = ""
	config.Protocol = protocol
	config.SrcAddr = strings.Split(fromAddr, ":")[0]
	config.SrcPort = Utils.ToInt(strings.Split(fromAddr, ":")[1])
	config.DestAddr = strings.Split(toAddr, ":")[0]
	config.DestPort = Utils.ToInt(strings.Split(toAddr, ":")[1])
	config.Status = 0
	Service.ForWardServ.CloseForward(config)

	c.Data["json"] = Models.FuncResult{Code: 0, Msg: ""}

	c.ServeJSON()
}
