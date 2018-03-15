package main

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/controllers"
	"github.com/csuhan/csugo/models"
	_ "github.com/csuhan/csugo/routers"
)

func main() {
	//日志
	beego.SetLogger("file", `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.ErrorController(&controllers.ErrorController{})
	models.InitDB() //初始化数据库连接
	beego.Run()
}
