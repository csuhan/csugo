package controllers

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
	"strings"
)

type LibController struct {
	beego.Controller
}

// @router /lib/login/:id/:pwd [get]
func (this *LibController) Login() {
	lib := &models.Lib{}
	_, err := lib.Login(this.Ctx.Input.Param(":id"), this.Ctx.Input.Param(":pwd"))
	if err != nil { //登录失败
		this.Data["json"] = struct {
			StateCode int
			Error     string
		}{
			StateCode: -1,
			Error:     err.Error(),
		}
		this.ServeJSON()
	}
	//登录成功
	this.Data["json"] = struct {
		StateCode int
		Error     string
	}{
		StateCode: 1,
		Error:     "",
	}
	this.ServeJSON()
}

// @router /lib/list/:id/:pwd [get]
func (this *LibController) List() {
	lib := &models.Lib{}
	books, err := lib.List(this.Ctx.Input.Param(":id"), this.Ctx.Input.Param(":pwd"))
	if err != nil {
		this.Data["json"] = struct {
			StateCode int
			Error     string
		}{
			StateCode: -1,
			Error:     err.Error(),
		}
		this.ServeJSON()
	}
	//登录成功
	this.Data["json"] = struct {
		StateCode int
		Error     string
		Books     []models.Book
	}{
		StateCode: 1,
		Error:     "",
		Books:     books,
	}
	this.ServeJSON()
}

// @router /lib/reloan/:id/:pwd/:books [get]
func (this *LibController) Reloan() {
	bookBarCodes := strings.Split(this.Ctx.Input.Param(":books"), "+")
	lib := &models.Lib{}
	books, err := lib.Borrow(this.Ctx.Input.Param(":id"), this.Ctx.Input.Param(":pwd"), bookBarCodes)
	if err != nil {
		this.Data["json"] = struct {
			StateCode int
			Error     string
		}{
			StateCode: -1,
			Error:     err.Error(),
		}
		this.ServeJSON()
	}

	this.Data["json"] = struct {
		StateCode int
		Error     string
		Books     []models.Book
	}{
		StateCode: 1,
		Error:     "",
		Books:     books,
	}
	this.ServeJSON()
}
