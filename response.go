package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liudeihao/furring/log"
)

func LogError(c *gin.Context, err error) {
	log.Error(fmt.Sprintf("error=%v, method=%v, url=%v", err, c.Request.Method, c.Request.URL))
}

func InternalServerErrorResponse(c *gin.Context, err error) {
	LogError(c, err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "服务器内部错误",
	})
}

func BadRequestResponse(c *gin.Context, msg ...any) {
	resp := gin.H{
		"error": "请求参数有误",
	}
	if len(msg) == 1 {
		if v, ok := msg[0].(string); ok {
			resp["details"] = v
		}
	}
	c.JSON(http.StatusBadRequest, resp)
}

func NotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "资源不存在",
	})
}
