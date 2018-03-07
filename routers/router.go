package routers

import (
	"github.com/csuhan/csugo/controllers"
	"github.com/astaxie/beego"
	_"github.com/csuhan/csugo/middleware"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    ns:=beego.NewNamespace("/api",
			beego.NSNamespace("/v1",
				beego.NSInclude(&controllers.JwcController{}),
				beego.NSInclude(&controllers.BusController{}),
				beego.NSInclude(&controllers.JobController{}),
			),
    	)

    beego.AddNamespace(ns)
}
