package handler

import (
    "errors"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/middleware"
    "github.com/liudeihao/furring/model"
    "github.com/liudeihao/furring/pkg/response"
    "github.com/liudeihao/furring/service"
)

type PostHandler struct {
    postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
    return &PostHandler{postService: postService}
}

func (h *PostHandler) RegisterRoutes(r gin.IRouter) {
    post := r.Group("/post")

    post.GET("/:id", middleware.MustParseID(), h.Get)

    private := r.Group("/post")
    private.Use(middleware.JWTAuthMiddleware())

    private.POST("/", middleware.MustParseID(), h.Post)
    private.PATCH("/:id", middleware.MustParseID(), h.Edit)
    private.DELETE("/:id", middleware.MustParseID(), h.Delete)
}

func (h *PostHandler) Get(c *gin.Context) {
    id := ParseID(c)

    r, err := h.postService.GetByID(id)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrPostNotFound):
            response.NotFound(c, "获取帖子失败")
        default:
            response.Internal(c, err)
        }
        return
    }
    response.OK(c, gin.H{
        "post": r,
    })
}

func (h *PostHandler) Post(c *gin.Context) {
    userid := ParseUID(c)

    var req model.PostCreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, err)
        return
    }
    req.UserID = userid
    resp, err := h.postService.Create(&req)
    if err != nil {
        response.Internal(c, err)
        return
    }
    response.OK(c, gin.H{
        "message": "发帖成功",
        "post":    resp,
    })
}

func (h *PostHandler) Edit(c *gin.Context) {
    id := ParseID(c)
    uid := ParseUID(c)
    var req model.PostUpdateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, err)
        return
    }
    req.PostID = id
    err := h.postService.Update(uid, req)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrPostNotFound):
            response.NotFound(c, "编辑帖子失败")
        case errors.Is(err, service.ErrNotAuthorized):
            response.Unauthorized(c, "您没有权限修改其他用户的帖子。")
        default:
            response.Internal(c, err)
        }
        return
    }
    response.OK(c, gin.H{
        "message": "修改成功",
    })
}

func (h *PostHandler) Delete(c *gin.Context) {
    id := ParseID(c)
    uid := ParseUID(c)
    err := h.postService.Delete(uid, id)

    if err != nil {
        switch {
        case errors.Is(err, service.ErrPostNotFound):
            response.NotFound(c, "删除帖子失败")
        case errors.Is(err, service.ErrNotAuthorized):
            response.Unauthorized(c, "您没有权限删除其他用户的帖子。")
        }
        return
    }
    response.OK(c, gin.H{
        "message": "删除成功",
    })
}
