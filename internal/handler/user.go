package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/internal/model"
    "github.com/liudeihao/furring/internal/service"
)

type UserHandler struct {
    userService *service.UserService
    postService *service.PostService
}

func NewUserHandler(userService *service.UserService, postService *service.PostService) *UserHandler {
    return &UserHandler{
        userService: userService,
        postService: postService,
    }
}

func (h *UserHandler) RegisterHandlers(r *gin.RouterGroup) {
    r.GET("/register", nil)
    r.GET("/login", nil)
}

func (h *UserHandler) Login(c *gin.Context) {

}

func (h *UserHandler) Logout(c *gin.Context) {

}

func (h *UserHandler) Register(c *gin.Context) {
    user := &model.User{}
    if err := c.ShouldBindJSON(user); err != nil {
        ErrorResponse(c, ErrBadRequest)
        return
    }
    user, err := h.userService.CreateUser(user)
    if err != nil {
        ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "user":    user,
        "message": "注册成功！",
    })
}
