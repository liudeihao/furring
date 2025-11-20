package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/pkg/contextkey"
)

type Handler interface {
    RegisterRoutes(r gin.IRouter)
}

func ParseID(c *gin.Context) uint {
    id, ok := c.Get(contextkey.QueryID)
    if !ok {
        panic("ParseID出现致命错误: 没有收到合法的id")
    }
    return id.(uint)
}

func ParseUID(c *gin.Context) uint {
    uid, ok := c.Get(contextkey.UserID)
    if !ok {
        panic("ParseUID出现致命错误: 没有收到合法的user_id")
    }
    return uid.(uint)
}
