package controllers

import (
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/models"
)

type NewsController struct {
	beego.Controller
}

// @router /news/list/:id
func (this *NewsController) GetNewsList() {
	pageid := this.Ctx.Input.Param(":id")
	news, err := models.GetNewsList(pageid)
	stateCode := 1
	errorstr := ""
	if err != nil {
		stateCode = -1
		errorstr = err.Error()
	}
	this.Data["json"] = struct {
		StateCode int
		Error     string
		News      models.NewsList
	}{
		StateCode: stateCode,
		Error:     errorstr,
		News:      news,
	}
	this.ServeJSON()
}

// @router /news/article/:link [get]
func (this *NewsController) GetNewsContent() {
	content, err := models.GetNewsContent(this.Ctx.Input.Param(":link"))
	stateCode := 1
	errorstr := ""
	if err != nil {
		stateCode = -1
		errorstr = err.Error()
	}
	this.Data["json"] = struct {
		StateCode int
		Error     string
		Content   string
	}{
		StateCode: stateCode,
		Error:     errorstr,
		Content:   content,
	}
	this.ServeJSON()
}
