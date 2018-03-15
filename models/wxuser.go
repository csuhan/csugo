package models

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/csuhan/csugo/utils"
)

type WxUser struct {
	Code       string `json:"code"`
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	WxToken    string `json:"wxtoken"`
	SchoolID   string `json:"schoolid"`
	JwcPwd     string `json:"jwcpwd"` //教务处
	ItsPwd     string `json:"itspwd"` //信息门户
	LibPwd     string `json:"libpwd"` //图书馆
}

//插入用户信息
func (this *WxUser) Insert() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	data, err := json.Marshal(this)
	if err != nil {
		return err
	}
	db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("wxuser")).Put([]byte(this.WxToken), data)
	})
	return nil
}

//获取用户信息
func (this *WxUser) Get() error {
	db, err := GetDB()
	if err != nil {
		return utils.ERROR_SERVER
	}
	defer db.Close()
	var value []byte
	db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket([]byte("wxuser")).Get([]byte(this.OpenID))
		return nil
	})
	//用户不存在
	if json.Unmarshal(value, this) != nil {
		return utils.ERROR_NO_USER
	}
	return nil
}
