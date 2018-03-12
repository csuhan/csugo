package main

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/controllers"
	_ "github.com/csuhan/csugo/routers"
)

func main() {
	//日志
	beego.SetLogger("file", `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
