package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["forward-server/Controllers:DefaultCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:DefaultCtrl"],
		beego.ControllerComments{
			Method:           "Default",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:DefaultCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:DefaultCtrl"],
		beego.ControllerComments{
			Method:           "ApiAuthFail",
			Router:           `/apiAuthFail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "AddForward",
			Router:           `/u/AddForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "ApiDoc",
			Router:           `/u/ApiDoc`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "ClearNetAgentStatus",
			Router:           `/u/ClearNetAgentStatus`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "OpenForward",
			Router:           `/u/OpenForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "CloseForward",
			Router:           `/u/CloseForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "OpenAllForward",
			Router:           `/u/OpenAllForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "CloseAllForward",
			Router:           `/u/CloseAllForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "ChangeForwardDebug",
			Router:           `/u/ChangeForwardDebug`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "CloseMagicService",
			Router:           `/u/CloseMagicService`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "DelForward",
			Router:           `/u/DelForward`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "EditForward",
			Router:           `/u/EditForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "ForwardList",
			Router:           `/u/ForwardList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "ForwardListJson",
			Router:           `/u/ForwardList/json`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "GetMagicStatus",
			Router:           `/u/GetMagicStatus`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "GetNetAgentStatus",
			Router:           `/u/GetNetAgentStatus`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "NetAgent",
			Router:           `/u/NetAgent`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "OpenMagicService",
			Router:           `/u/OpenMagicService`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "SaveForward",
			Router:           `/u/SaveForward`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "StartAgentJob",
			Router:           `/u/StartAgentJob`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "StopAgentJob",
			Router:           `/u/StopAgentJob`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "AddBatchForward",
			Router:           `/u/AddBatchForward`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "SaveBatchForward",
			Router:           `/u/SaveBatchForward`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "ImportForward",
			Router:           `/u/ImportForward`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method:           "SaveImportForward",
			Router:           `/u/SaveImportForward`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:LoginCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:LoginCtrl"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:LoginCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:LoginCtrl"],
		beego.ControllerComments{
			Method:           "DoLogin",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:LoginCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:LoginCtrl"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method:           "CloseForward",
			Router:           `/CloseForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method:           "OpenForward",
			Router:           `/OpenForward`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method:           "ServerSummary",
			Router:           `/ServerSummary`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method:           "ChangePwd",
			Router:           `/u/changePwd`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method:           "DoChangePwd",
			Router:           `/u/doChangePwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method:           "GetServerTime",
			Router:           `/u/getServerTime`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/u/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["forward-server/Controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method:           "Main",
			Router:           `/u/main`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
