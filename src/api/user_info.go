package api

import (
	"0offer/model"
	"0offer/service"
	"0offer/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func SelectAllUserInfo(c *gin.Context) {
	if userInfos, err := service.SelectAllUserInfo(); err == nil {
		utils.OkWithData(userInfos, c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

func BatchInsertUserInfo(c *gin.Context) {
	var body struct {
		UserInfos []model.UserInfo
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Fatalf("AddUserInfos failed with %s\n", err)
		return
	}
	if userInfos, err := service.BatchInsertUserInfo(body.UserInfos); err == nil {
		utils.OkWithData(userInfos, c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
