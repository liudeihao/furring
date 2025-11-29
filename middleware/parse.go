package middleware

import (
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/pkg/contextkey"
    "github.com/liudeihao/furring/pkg/response"
)

func MustParseID() gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            response.NotFound(c, "资源获取失败")
            return
        }
        c.Set(contextkey.QueryID, uint(id))
        c.Next()
    }
}
