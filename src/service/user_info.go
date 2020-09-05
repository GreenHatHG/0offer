package service

import (
	"0offer/global"
	"0offer/model"
	"github.com/gin-gonic/gin"
)

func SelectAllUserInfo() map[string]interface{} {
	var userInfo []model.UserInfo
	_ = global.DB.Find(&userInfo)
	return gin.H{
		"userInfos": userInfo,
	}
}
