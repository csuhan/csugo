package controllers

import(
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
)

type JwcController struct{
	beego.Controller
}

// @router /jwc/:id/:pwd/grade [get]
func (this *JwcController)Grade(){
	user:=&models.JwcUser{
		Id:this.Ctx.Input.Param(":id"),
		Pwd:this.Ctx.Input.Param(":pwd")}
	jwc:=&models.Jwc{}
	grade,err:=jwc.Grade(user)
	stateCode:=1
	errorstr:=""
	if err!=nil{
		stateCode=-1
		errorstr=err.Error()
	}
	this.Data["json"]= struct {
		StateCode int
		Error string
		Grades []models.JwcGrade
	}{
		StateCode:stateCode,
		Error:errorstr,
		Grades:grade,
	}
	this.ServeJSON()
}

// @router /jwc/:id/:pwd/rank [get]
func (this *JwcController)Rank(){
	user:=&models.JwcUser{
		Id:this.Ctx.Input.Param(":id"),
		Pwd:this.Ctx.Input.Param(":pwd")}
	jwc:=&models.Jwc{}
	rank,err:=jwc.Rank(user)
	stateCode:=1
	errorstr:=""
	if err!=nil{
		stateCode=-1
		errorstr=err.Error()
	}
	this.Data["json"]= struct {
		StateCode int
		Error string
		Rank []map[string]models.Rank
	}{
		StateCode:stateCode,
		Error:errorstr,
		Rank:rank,
	}
	this.ServeJSON()
}

// @router /jwc/:id/:pwd/class/:term/:week [get]
func (this *JwcController)Class(){
	user:=&models.JwcUser{
		Id:this.Ctx.Input.Param(":id"),
		Pwd:this.Ctx.Input.Param(":pwd")}
	week:=this.Ctx.Input.Param(":week")
	term:=this.Ctx.Input.Param(":term")
	jwc:=&models.Jwc{}
	class,err:=jwc.Class(user,week,term)
	stateCode:=1
	errorstr:=""
	if err!=nil{
		stateCode=-1
		errorstr=err.Error()
	}
	this.Data["json"]= struct {
		StateCode int
		Error string
		Class [][]models.Class
	}{
		StateCode:stateCode,
		Error:errorstr,
		Class:class,
	}
	this.ServeJSON()
}