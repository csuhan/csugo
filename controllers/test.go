package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {

	this.Data["json"] = "s"
	this.ServeJSON()
}
func (this *TestController) Post() {

	this.Data["json"] = "ss"
	this.ServeJSON()
}
