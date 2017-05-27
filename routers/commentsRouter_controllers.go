package routers

import (
	"github.com/astaxie/beego"
)

func init() {

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

	beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:LoginCtrl"],
		beego.ControllerComments{
			Method: "ApiAuthFail",
			Router: `/apiAuthFail`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method: "ServerSummary",
			Router: `/ServerSummary`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method: "OpenTcpForward",
			Router: `/OpenTcpForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:RestApiCtrl"],
		beego.ControllerComments{
			Method: "CloseTcpForward",
			Router: `/CloseTcpForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:TcpForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:TcpForwardCtrl"],
		beego.ControllerComments{
			Method: "TcpForwardList",
			Router: `/u/TcpForwardList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:TcpForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:TcpForwardCtrl"],
		beego.ControllerComments{
			Method: "OpenTcpForward",
			Router: `/u/OpenTcpForward`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["port-forward/controllers:TcpForwardCtrl"] = append(beego.GlobalControllerRouter["port-forward/controllers:TcpForwardCtrl"],
		beego.ControllerComments{
			Method: "CloseTcpForward",
			Router: `/u/CloseTcpForward`,
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

}
