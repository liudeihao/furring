package service

import (
    "time"

    jwt "github.com/golang-jwt/jwt/v4"
    "github.com/liudeihao/furring/pkg/errors"
)

type JWTService struct {
    secretKey string
    issuer    string
}

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.RegisteredClaims
}

func NewJWTService(secretKey string, issuer string) *JWTService {
    return &JWTService{secretKey, issuer}
}

func (s *JWTService) GenerateToken(userID uint) (string, error) {
    expirationTime := time.Now().Add(7 * 24 * time.Hour)

    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer:    s.issuer,
            Subject:   "user_token",
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("验证Token时，遇到unexpected signing method")
        }
        return []byte(s.secretKey), nil
    })
}

func (s *JWTService) ExtractUserID(tokenString string) (uint, error) {
    token, err := s.ValidateToken(tokenString)
    if err != nil {
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // 注意：jwt.MapClaims中的数字默认是float64类型
        if userID, exists := claims["user_id"]; exists {
            return uint(userID.(float64)), nil
        }
        return 0, errors.New("user_id not found in token")
    }

    return 0, errors.New("invalid token claims")
}
