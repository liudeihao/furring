package handler

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/pkg/errors"
    "github.com/liudeihao/furring/pkg/log"
)

var (
    ErrBadID      = errors.New("错误的ID")
    ErrBadRequest = errors.New("请求参数有误")
)

func LogError(c *gin.Context, err error) {
    log.Error(fmt.Sprintf("error=%v, method=%v, url=%v", err, c.Request.Method, c.Request.URL))
}

func ErrorResponse(c *gin.Context, err error) {
    if errors.IsInternalError(err) {
        LogError(c, err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "服务器内部错误",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "error": err.Error(),
    })
}
