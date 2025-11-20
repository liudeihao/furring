package handler

import (
    "errors"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/middleware"
    "github.com/liudeihao/furring/model"
    "github.com/liudeihao/furring/pkg/response"
    "github.com/liudeihao/furring/service"
)

type CommentHandler struct {
    commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
    return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) RegisterRoutes(router gin.IRouter) {
    comment := router.Group("/comment")
    comment.POST("/", h.Comment)

    private := comment.Group("/comment")
    private.Use(middleware.JWTAuthMiddleware())

    private.DELETE("/", h.DeleteComment)
}

func (h *CommentHandler) GetComment(c *gin.Context) {

}

func (h *CommentHandler) Comment(c *gin.Context) {
    uid := ParseUID(c)
    var comment model.CommentCreateRequest
    if err := c.ShouldBindJSON(&comment); err != nil {
        response.BadRequest(c, err)
        return
    }
    comment.UserID = uid
    cid, err := h.commentService.Comment(comment)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrCommentNotFound):
            response.NotFound(c, "要评论的帖子不存在")
        default:
            response.Internal(c, err)
        }
    }
    response.OK(c, gin.H{
        "message": "评论成功",
        "cid":     cid,
    })
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
    uid := ParseUID(c)
    id := ParseID(c)
    err := h.commentService.Delete(uid, id)
    if err != nil {
        switch {
        case errors.Is(err, service.ErrCommentNotFound):
            response.NotFound(c, "要删除的评论不存在")
        case errors.Is(err, service.ErrNotAuthorized):
            response.Unauthorized(c, "您没有权限删除其他用户的评论。")
        default:
            response.Internal(c, err)
        }
    }
    response.OK(c, gin.H{
        "message": "评论删除成功",
    })

}
