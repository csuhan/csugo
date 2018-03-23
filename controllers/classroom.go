package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
)

type ClassRoomController struct {
	beego.Controller
}

// @router /classroom/:term/:week/:xq/:jxl [get]
func (this *ClassRoomController) GetFreeWeekTime() {
	term := this.Ctx.Input.Param(":term")
	week := this.Ctx.Input.Param(":week")
	xq := this.Ctx.Input.Param(":xq")
	jxl := this.Ctx.Input.Param(":jxl")

	cls, err := models.GetFreeWeekTime(term, week, xq, jxl)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["json"] = cls
	this.ServeJSON()

}
