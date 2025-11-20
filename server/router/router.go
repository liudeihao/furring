package router

import (
    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/handler"
)

func New(handlers []handler.Handler) *gin.Engine {
    r := gin.Default()
    api := r.Group("/api")

    for _, h := range handlers {
        h.RegisterRoutes(api)
    }

    return r
}
