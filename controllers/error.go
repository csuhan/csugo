package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

type Error struct {
	StateCode int
	Error     string
}

func (this *ErrorController) Error404() {
	this.Data["json"] = Error{
		StateCode: 404,
		Error:     "api not found",
	}
	this.TplName = "errors/404.html"
	this.ServeJSON()
}
