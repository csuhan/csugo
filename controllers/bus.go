package controllers

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
)

type BusController struct {
	beego.Controller
}

// @router /bus/search/:start/:end/:time [get]
func (this *BusController) Search() {
	start := this.Ctx.Input.Param(":start")
	end := this.Ctx.Input.Param(":end")
	time := this.Ctx.Input.Param(":time")
	bus := new(models.Bus)
	buses, err := bus.Search(start, end, time)
	stateCode := 1
	errorstr := ""
	if err != nil {
		stateCode = -1
		errorstr = err.Error()
	}
	this.Data["json"] = struct {
		StateCode int
		Error     string
		Buses     []models.Bus
	}{
		StateCode: stateCode,
		Error:     errorstr,
		Buses:     buses,
	}
	this.ServeJSON()
}
