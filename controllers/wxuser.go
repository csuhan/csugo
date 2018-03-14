package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/csuhan/csugo/utils"
	"github.com/astaxie/beego/httplib"
)

type WxUser struct{
	Code string  `json:"code"`
	OpenID string `json:"openid"`
	ExpiresIn int `json:"expires_in"`
	SessionKey string `json:"session_key"`
}
type Status struct{
	StateCode int
	Error string
}

type WxUserController struct{
	beego.Controller
}

// @router /login [post]
func (this *WxUserController)Login(){
	//解析code
	var user WxUser
	if err:=json.Unmarshal(this.Ctx.Input.RequestBody,&user);err!=nil{
		this.Data["json"]=&Status{
			StateCode:-1,
			Error:utils.ERROR_DATA.Error(),
		}
		this.ServeJSON()
	}
	appID:=beego.AppConfig.String("AppID")
	appSecret:=beego.AppConfig.String("AppSecret")
	//用code换取openid,seesion_key
	req:=httplib.Get("https://api.weixin.qq.com/sns/jscode2session?appid="+appID+"&secret="+appSecret+"&js_code="+user.Code+"&grant_type=authorization_code")
	if err:=req.ToJSON(&user);err!=nil{
		this.Data["json"]=&Status{
			StateCode:-1,
			Error:utils.ERROR_SERVER.Error(),
		}
		this.ServeJSON()
	}
	this.Data["json"]=user
	this.ServeJSON()
}
