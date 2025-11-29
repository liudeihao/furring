package middleware

import (
    "fmt"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "github.com/liudeihao/furring/config"
    "github.com/liudeihao/furring/pkg/contextkey"
    "github.com/liudeihao/furring/pkg/response"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            response.Unauthorized(c, "Authorization为空")
            return
        }
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            response.Unauthorized(c, "Authorization格式错误")
            return
        }
        tokenString := parts[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return []byte(config.Instance.JWT.SecretKey), nil
        })

        if err != nil || !token.Valid {
            fmt.Println(err, token.Valid)
            response.Unauthorized(c, "token无效")
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            response.Unauthorized(c, "token claim无效")
            return
        }

        // claim存入Context
        claimUserID, exists := claims["user_id"]
        if !exists {
            response.Unauthorized(c, "token中不存在user_id")
            return
        }

        if fID, ok := claimUserID.(float64); ok {
            c.Set(contextkey.UserID, uint(fID))
        } else {
            c.Set(contextkey.UserID, claimUserID)
        }

        c.Next()
    }
}
