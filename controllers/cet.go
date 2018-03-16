package controllers

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
)

type CetController struct{
	beego.Controller
}

// @router /cet/zkz/:id/:type [get]
func (this *CetController)GetZKZ(){
	cet:=&models.Cet{}
	zkz,err:=cet.GetZKZ(this.Ctx.Input.Param(":id"),this.Ctx.Input.Param(":type"))
	if err!=nil{
		this.Data["json"]= struct {
			StateCode int
			Error string
		}{
			StateCode:-1,
			Error:err.Error(),
		}
		this.ServeJSON()
	}
	this.Data["json"]=struct {
		StateCode int
		Error string
		ZKZ models.ZKZH
	}{
		StateCode:1,
		Error:"",
		ZKZ: zkz,
	}
	this.ServeJSON()
}

// @router /cet/hgrade/:id/:name [get]
func (this *CetController)GetHGrade(){
	cet:=&models.Cet{}
	hgrade,err:=cet.GetHGrade(this.Ctx.Input.Param(":id"),this.Ctx.Input.Param(":name"))
	if err!=nil{
		this.Data["json"]= struct {
			StateCode int
			Error string
		}{
			StateCode:-1,
			Error:err.Error(),
		}
		this.ServeJSON()
	}
	this.Data["json"]=struct {
		StateCode int
		Error string
		HGrades []models.HGrade
	}{
		StateCode:1,
		Error:"",
		HGrades:hgrade,
	}
	this.ServeJSON()
}