package service

import (
    "time"

    "github.com/golang-jwt/jwt/v4"
    "github.com/liudeihao/furring/config"
)

func GenerateToken(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "iss":     config.Instance.JWT.Issuer,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // 这里必须是[]byte，不能是string
    return token.SignedString([]byte(config.Instance.JWT.SecretKey))
}
