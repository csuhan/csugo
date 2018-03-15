package routers

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/controllers"
	_ "github.com/csuhan/csugo/middleware"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(&controllers.JwcController{}),
			beego.NSInclude(&controllers.BusController{}),
			beego.NSInclude(&controllers.JobController{}),
		),
	)
	wx := beego.NewNamespace("/wxapp",
		beego.NSInclude(&controllers.WxUserController{}),
	)
	beego.AddNamespace(wx)
	beego.AddNamespace(ns)
}
