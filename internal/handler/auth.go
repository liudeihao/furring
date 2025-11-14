package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/internal/model"
    "github.com/liudeihao/furring/internal/service"
)

type AuthHandler struct {
    authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterHandlers(r *gin.RouterGroup) {
    r.GET("/login", nil)
    r.GET("/logout", nil)
    r.GET("/register", nil)
}

func (h *AuthHandler) Register(c *gin.Context) {
    req := new(model.RegisterRequest)
    if err := c.ShouldBindJSON(&req); err != nil {
        ErrorResponse(c, ErrBadRequest)
        return
    }
    user, err := h.authService.Register(req)
    if err != nil {
        ErrorResponse(c, err)
        return
    }
    token, err := h.authService.GenerateToken(user)
    if err != nil {
        ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "注册成功",
        "user":    user,
        "token":   token,
    })
}

func (h *AuthHandler) Login(c *gin.Context) {
    req := new(model.LoginRequest)
    if err := c.ShouldBindJSON(req); err != nil {
        ErrorResponse(c, ErrBadRequest)
        return
    }
    user, token, err := h.authService.Login(req)
    if err != nil {
        ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "登录",
        "user":    user,
        "token":   token,
    })
}

func (h *AuthHandler) Logout(c *gin.Context) {

}
