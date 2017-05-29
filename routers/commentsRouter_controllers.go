package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["port-forward/controllers:DefaultCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:DefaultCtrl"],
		beego.ControllerComments{
			Method: "Default",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:DefaultCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:DefaultCtrl"],
		beego.ControllerComments{
			Method: "ApiAuthFail",
			Router: `/apiAuthFail`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "ForwardList",
			Router: `/u/ForwardList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "ForwardListJson",
			Router: `/u/ForwardList/json`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "AddForward",
			Router: `/u/AddForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "EditForward",
			Router: `/u/EditForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "DelForward",
			Router: `/u/DelForward`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "SaveForward",
			Router: `/u/SaveForward`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "OpenForward",
			Router: `/u/OpenForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "CloseForward",
			Router: `/u/CloseForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:ForwardCtrl"],
		beego.ControllerComments{
			Method: "ApiDoc",
			Router: `/u/ApiDoc`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:HelpCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:HelpCtrl"],
		beego.ControllerComments{
			Method: "GetTcp",
			Router: `/GetTcp`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"],
		beego.ControllerComments{
			Method: "DoLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method: "ServerSummary",
			Router: `/ServerSummary`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method: "OpenForward",
			Router: `/OpenForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method: "CloseForward",
			Router: `/CloseForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method: "Main",
			Router: `/u/main`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/u/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:UCenterCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:UCenterCtrl"],
		beego.ControllerComments{
			Method: "GetServerTime",
			Router: `/u/getServerTime`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
