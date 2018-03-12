package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:BusController"] = append(beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:BusController"],
		beego.ControllerComments{
			Method:           "Search",
			Router:           `/bus/search/:start/:end/:time`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JobController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/job/:typeid/:pageindex/:pagesize/:hastime`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JwcController"] = append(beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JwcController"],
		beego.ControllerComments{
			Method:           "Class",
			Router:           `/jwc/:id/:pwd/class/:term/:week`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JwcController"] = append(beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JwcController"],
		beego.ControllerComments{
			Method:           "Grade",
			Router:           `/jwc/:id/:pwd/grade`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JwcController"] = append(beego.GlobalControllerRouter["github.com/csuhan/csugo/controllers:JwcController"],
		beego.ControllerComments{
			Method:           "Rank",
			Router:           `/jwc/:id/:pwd/rank`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
