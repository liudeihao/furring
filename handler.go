package main

import (
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/log"
    "gorm.io/gorm"
)

type Handler struct {
    db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
    return &Handler{db: db}
}

func (h *Handler) RegisterHandlers(r *gin.Engine) {
    users := r.Group("/users")
    users.GET("", h.GetUsersAll)
    users.POST("/", h.CreateUser)
    users.GET("/:id", h.GetUserByID)
    users.DELETE("/:id", h.DeleteUser)
    users.GET("/:id/posts", h.GetUserPosts)

    posts := r.Group("/posts")
    posts.GET("", h.GetPostsAll)
    posts.POST("/", h.CreatePost)
    posts.GET("/:id", h.GetPostByID)
    posts.DELETE("/:id", h.DeletePost)
}

func (h *Handler) GetUsersAll(c *gin.Context) {
    var users []User
    if err := h.db.Find(&users).Error; err != nil {
        InternalServerErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "users": users,
    })
}

func (h *Handler) GetUserByID(c *gin.Context) {
    var user User
    id := c.Param("id")
    if err := h.db.First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            NotFoundResponse(c)
            return
        } else {
            InternalServerErrorResponse(c, err)
            return
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}

func (h *Handler) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    result := h.db.Delete(&User{}, id)
    if result.Error != nil {
        InternalServerErrorResponse(c, result.Error)
        return
    }
    if result.RowsAffected == 0 {
        NotFoundResponse(c)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "删除成功",
    })
}

func (h *Handler) CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        BadRequestResponse(c)
        return
    }
    if err := h.db.Create(&user).Error; err != nil {
        InternalServerErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "user":    user,
        "message": "创建成功",
    })
}

func (h *Handler) GetPostsAll(c *gin.Context) {
    var posts []Post
    if err := h.db.Find(&posts).Error; err != nil {
        InternalServerErrorResponse(c, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "posts": posts,
    })
}

func (h *Handler) GetPostByID(c *gin.Context) {
    var post Post
    id := c.Param("id")
    if err := h.db.First(&post, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            NotFoundResponse(c)
            return
        } else {
            InternalServerErrorResponse(c, err)
            return
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "post": post,
    })
}

func (h *Handler) DeletePost(c *gin.Context) {
    id := c.Param("id")
    result := h.db.Delete(&Post{}, id)
    if result.Error != nil {
        InternalServerErrorResponse(c, result.Error)
        return
    }
    if result.RowsAffected == 0 {
        NotFoundResponse(c)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "删除成功",
    })
}

func (h *Handler) CreatePost(c *gin.Context) {
    var post Post
    if err := c.ShouldBindJSON(&post); err != nil {
        BadRequestResponse(c)
        return
    }
    log.Debugf("%#v", post)
    userid := post.UserID
    if !h.userExists(userid) {
        BadRequestResponse(c, "用户不存在")
        return
    }

    if err := h.db.Create(&post).Error; err != nil {
        InternalServerErrorResponse(c, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "post":    post,
        "message": "创建成功",
    })
}

func (h *Handler) GetUserPosts(c *gin.Context) {
    userid := c.Param("id")
    var posts []Post
    if err := h.db.Where("user_id = ?", userid).Find(&posts).Error; err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            InternalServerErrorResponse(c, err)
            return
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "posts":   posts,
        "n_posts": len(posts),
    })
}

func (h *Handler) userExists(id int) bool {
    if err := h.db.First(&User{}, id).Error; err != nil {
        return false
    }
    return true
}
