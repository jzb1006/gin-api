package wechat

import (
	"github.com/medivhzhan/weapp/v2"
	"go-gin-api/config"
	"go-gin-api/init/mysql"
	"go-gin-api/model"
)

func Login(code string) (res *weapp.LoginResponse, err error)  {
	res, err = weapp.Login(config.GinConfig.WechatMP.AppId, config.GinConfig.WechatMP.AppsSecret, code)
	return
}

func Exist(openId string) (user model.User, err error) {
	err = mysql.DEFAULTDB.Where("openid = ?", openId).First(&user).Error
	return
}

func Save(openid string, sessionKey string, unionid string) (user model.User, err error)  {
	user = model.User{Openid:openid,SessionKey:sessionKey,Unionid:unionid}
	err = mysql.DEFAULTDB.Create(&user).Error
	return
}

//func Update(rawData string)(user model.User, err error)  {
//
//}