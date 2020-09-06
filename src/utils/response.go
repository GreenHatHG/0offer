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

func Ok(c *gin.Context) {
	Result(http.StatusOK, 200, map[string]interface{}{}, "操作成功", c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, 0, data, "操作成功", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusBadRequest, 400, map[string]interface{}{}, message, c)
}
