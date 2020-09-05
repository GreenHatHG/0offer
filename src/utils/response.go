package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(httpCode int, code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(httpCode, Response{
		code,
		data,
		msg,
	})
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, 0, data, "操作成功", c)
}
