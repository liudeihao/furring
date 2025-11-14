package handler

import "github.com/gin-gonic/gin"

type Handler interface {
    RegisterHandlers(group *gin.RouterGroup)
}
