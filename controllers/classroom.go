package controllers

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
)

type ClassRoomController struct {
	beego.Controller
}

// @router /classroom/time/:term/:week/:xq/:jxl [get]
func (this *ClassRoomController) GetFreeWeekTime() {
	term := this.Ctx.Input.Param(":term")
	week := this.Ctx.Input.Param(":week")
	xq := this.Ctx.Input.Param(":xq")
	jxl := this.Ctx.Input.Param(":jxl")

	cls, err := models.GetFreeWeekTime(term, week, xq, jxl)
	stateCode := 1
	errorstr := ""
	if err != nil {
		stateCode = -1
		errorstr = err.Error()
	}
	this.Data["json"] = struct {
		StateCode int
		Error     string
		CLS       []models.ClassRoom
	}{
		StateCode: stateCode,
		Error:     errorstr,
		CLS:       cls,
	}
	this.ServeJSON()

}

// @router /classroom/jxl/:xq [get]
func (this *ClassRoomController) GetJXL() {
	xq := this.Ctx.Input.Param(":xq")
	jxls, err := models.GetBuildingsByXQ(xq)
	stateCode := 1
	errorstr := ""
	if err != nil {
		stateCode = -1
		errorstr = err.Error()
	}
	this.Data["json"] = struct {
		StateCode int
		Error     string
		JXLS      []models.JXL
	}{
		StateCode: stateCode,
		Error:     errorstr,
		JXLS:      jxls,
	}
	this.ServeJSON()
}

// @router /classroom/jxls [get]
func (this *ClassRoomController) GetJXLS() {
	jxls := models.GetBuildings()
	this.Data["json"] = struct {
		StateCode int
		Error     string
		JXLS      map[string][]models.JXL
	}{
		StateCode: 1,
		Error:     "",
		JXLS:      jxls,
	}
	this.ServeJSON()
}
