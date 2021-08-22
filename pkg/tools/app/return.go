package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code" example:"200"` // 代码
	Data    interface{} `json:"data"`               // 数据集
	Message string      `json:"message"`            // 消息
}

// Error 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	var res = Response{
		Code: code,
	}

	if err != nil && msg != "" {
		res.Message = fmt.Sprintf("%s, error: %s", msg, err.Error())
	} else if err != nil {
		res.Message = err.Error()
	} else if msg != "" {
		res.Message = msg
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	var res = Response{
		Code:    200,
		Data:    data,
		Message: "",
	}
	if msg != "" {
		res.Message = msg
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}
