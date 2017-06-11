package routers

import (
	"port-forward/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//
	beego.Include(&controllers.DefaultCtrl{})
	beego.Include(&controllers.LoginCtrl{})
	beego.Include(&controllers.UCenterCtrl{})
	beego.Include(&controllers.ForwardCtrl{})

	api_ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&controllers.RestApiCtrl{},
			),
		),
	)

	beego.AddNamespace(api_ns)

}
