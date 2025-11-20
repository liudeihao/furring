package handler

import (
    "errors"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/middleware"
    "github.com/liudeihao/furring/model"
    "github.com/liudeihao/furring/pkg/response"
    "github.com/liudeihao/furring/service"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterRoutes(r gin.IRouter) {
    user := r.Group("/user")

    user.POST("/login", h.Login)
    user.POST("/logout", h.Logout)
    user.POST("/register", h.Register)

    user.GET("/:id", middleware.MustParseID(), h.GetPublic)
    user.GET("/:id/posts", middleware.MustParseID(), h.GetPosts)

    private := user.Group("/private")
    private.Use(middleware.JWTAuthMiddleware())

    private.GET("/:id", middleware.MustParseID(), h.GetPrivate)
}

func (h *UserHandler) GetPosts(c *gin.Context) {
    id := ParseID(c)

    resp, err := h.userService.GetUserPosts(id)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrPostNotFound):
            response.NotFound(c, "无法获取帖子")
        default:
            response.Internal(c, err)
        }
        return
    }
    response.OK(c, gin.H{
        "user_id": id,
        "posts":   resp.Posts,
    })
}

func (h *UserHandler) GetPublic(c *gin.Context) {
    id := ParseID(c)

    info, err := h.userService.GetPublicInfoByID(id)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrUserNotFound):
            response.NotFound(c, "获取用户公开信息失败")
        default:
            response.Internal(c, err)
        }
        return
    }
    response.OK(c, gin.H{
        "user_public_info": info,
    })
}
func (h *UserHandler) GetPrivate(c *gin.Context) {
    id := ParseID(c)
    uid := ParseUID(c)

    info, err := h.userService.GetPrivateInfoByID(uid, id)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrUserNotFound):
            response.NotFound(c, "获取用户隐私信息失败")
        case errors.Is(err, service.ErrNotAuthorized):
            response.Unauthorized(c, "您没有权限查看其他用户的隐私信息。")
        default:
            response.Internal(c, err)
        }
        return
    }
    response.OK(c, gin.H{
        "user_private_info": info,
    })
}

type person struct {
    name string
    age  int
}

func (h *UserHandler) Login(c *gin.Context) {
    var r model.LoginRequest
    if err := c.ShouldBindJSON(&r); err != nil {
        response.BadRequest(c, err)
        return
    }
    resp, err := h.userService.Login(r)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrUserNotFound):
        case errors.Is(err, service.ErrPasswordWrong):
            response.Unauthorized(c, "登录失败, 密码或用户名错误")
        default:
            response.Internal(c, err)
        }
        return
    }
    response.OK(c, gin.H{
        "id":    resp.ID,
        "token": resp.Token,
    })

}

func (h *UserHandler) Logout(c *gin.Context) {
    // TODO：JWT黑名单
    response.OK(c, gin.H{
        "message": "退出成功",
    })
}

func (h *UserHandler) Register(c *gin.Context) {
    var r model.RegisterRequest
    if err := c.ShouldBindJSON(&r); err != nil {
        response.BadRequest(c, err)
        return
    }
    resp, err := h.userService.Register(r)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrUsernameDuplicate):
        case errors.Is(err, service.ErrEmailDuplicate):
            response.BadRequest(c, err)
        default:
            response.Internal(c, err)
            return
        }
    }
    response.OK(c, gin.H{
        "id":    resp.ID,
        "token": resp.Token,
    })
}
