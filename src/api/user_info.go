package api

import (
	"0offer/service"
	"0offer/utils"
	"github.com/gin-gonic/gin"
)

func SelectAllUserInfo(c *gin.Context) {
	utils.OkWithData(service.SelectAllUserInfo(), c)
}
