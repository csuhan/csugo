package controllers

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
)

type JobController struct{
	beego.Controller
}

// @router /job/:typeid/:pageindex/:pagesize/:hastime [get]
func (this *JobController)List(){
	params:=this.Ctx.Input.Params()
	typeid:=params[":typeid"]
	pagesize:=params[":pagesize"]
	pageindex:=params[":pageindex"]
	hastime:=params[":hastime"]
	job:=new(models.Job)
	jobs,err:=job.List(typeid,pageindex,pagesize,hastime)
	stateCode:=1
	errorstr:=""
	if err!=nil{
		stateCode=-1
		errorstr=err.Error()
	}
	this.Data["json"]= struct {
		StateCode int
		Error string
		Jobs []models.Job
	}{
		StateCode:stateCode,
		Error:errorstr,
		Jobs:jobs,
	}
	this.ServeJSON()

}
