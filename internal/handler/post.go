package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/internal/service"
)

type PostHandler struct {
    postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
    return &PostHandler{
        postService: postService,
    }
}

func (h *PostHandler) RegisterHandlers(r *gin.RouterGroup) {
    post := r.Group("/post")
    post.GET("/new", h.GetRecentPostList)
    post.GET("/hot", h.GetHotPostList)
    post.GET("/:id", h.GetPost)
}

func (h *PostHandler) GetPost(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ErrorResponse(c, ErrBadID)
        return
    }
    c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *PostHandler) GetRecentPostList(c *gin.Context) {

}

func (h *PostHandler) GetHotPostList(c *gin.Context) {

}
