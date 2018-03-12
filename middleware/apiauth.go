package middleware

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/csuhan/csugo/controllers"
)

//token认证
func init() {
	beego.InsertFilter("/api/*", beego.BeforeRouter, func(ctx *context.Context) {
		token := ctx.Input.Query("token")
		if token != "csugo-token" {
			data, _ := json.Marshal(controllers.Error{
				StateCode: 401,
				Error:     "unauthorized",
			})
			ctx.WriteString(string(data))
		}
	})
}
