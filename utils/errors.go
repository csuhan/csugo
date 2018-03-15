package utils

import "errors"

var (
	ERROR_SERVER  = errors.New("服务器出了点问题，重试一下？")
	ERROR_ID_PWD  = errors.New("账号或者密码错误,请重新输入")
	ERROR_UNKOWN  = errors.New("未知错误")
	ERROR_JWC     = errors.New("教务系统出了点问题,请重试")
	ERROR_DATA    = errors.New("数据错误")
	ERROR_NO_USER = errors.New("用户不存在")
)
