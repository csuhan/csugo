package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/csuhan/csugo/models"
	"github.com/csuhan/csugo/utils"
)

type Status struct {
	StateCode int
	Error     string
}

type WxUserController struct {
	beego.Controller
}

//用户登录
// @router /login [post]
func (this *WxUserController) Login() {
	//解析code
	var user models.WxUser
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err != nil {
		this.Data["json"] = &Status{
			StateCode: -1,
			Error:     utils.ERROR_DATA.Error(),
		}
		this.ServeJSON()
	}
	//用code换取openid,seesion_key
	appID := beego.AppConfig.String("AppID")
	appSecret := beego.AppConfig.String("AppSecret")
	req := httplib.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + appID + "&secret=" + appSecret + "&js_code=" + user.Code + "&grant_type=authorization_code")
	if err := req.ToJSON(&user); err != nil {
		this.Data["json"] = &Status{
			StateCode: -1,
			Error:     utils.ERROR_SERVER.Error(),
		}
		this.ServeJSON()
	}
	//生成wxtoken
	md5Token := md5.New()
	md5Token.Write([]byte(user.OpenID))
	user.WxToken = hex.EncodeToString(md5Token.Sum(nil))

	userTemp := user //临时复制对象

	//判断用户是否存在
	err := user.Get()
	//数据库错误
	if err == utils.ERROR_SERVER {
		this.Data["json"] = &Status{
			StateCode: -1,
			Error:     err.Error(),
		}
		this.ServeJSON()
	}
	//用户不存在,插入用户
	if err == utils.ERROR_NO_USER {
		user = userTemp
	}
	//用户存在,仅更新session_key
	if user.WxToken != "" {
		user.SessionKey = userTemp.SessionKey

	}
	//更新数据
	if err := user.Insert(); err != nil {
		this.Data["json"] = &Status{
			StateCode: -1,
			Error:     utils.ERROR_DATA.Error(),
		}
		this.ServeJSON()
	}
	//输出
	this.Data["json"] = struct {
		Wxtoken string
	}{
		Wxtoken: user.WxToken,
	}
	this.ServeJSON()
}
