package api

import (
	"0offer/service"
	"0offer/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func Miaosha(c *gin.Context) {
	var body struct {
		GoodsId uint
		UserId  uint
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Fatalf("Miaosha failed with %s\n", err)
		return
	}
	if err := service.Miaosha(body.GoodsId, body.UserId); err == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
