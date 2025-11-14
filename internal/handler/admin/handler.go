package admin

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/internal/handler"
    "github.com/liudeihao/furring/internal/model"
    "github.com/liudeihao/furring/internal/service"
    "github.com/liudeihao/furring/pkg/pagination"
)

type ManagementHandler struct {
    userService *service.UserService
    postService *service.PostService
}

func NewHandler(userService *service.UserService, postService *service.PostService) *ManagementHandler {
    return &ManagementHandler{
        userService: userService,
        postService: postService,
    }
}

func (h *ManagementHandler) RegisterHandlers(r *gin.RouterGroup) {
    admin := r.Group("/admin")
    users := admin.Group("/users")
    users.GET("", h.GetUsersAll)
    users.POST("/", h.CreateUser)
    users.PATCH("/:id", h.UpdateUser)
    users.GET("/:id", h.GetUser)
    users.DELETE("/:id", h.DeleteUser)

}

func (h *ManagementHandler) GetUsersAll(c *gin.Context) {
    p := pagination.New(c)

    users, err := h.userService.GetUserList(p.Offset(), p.Limit())
    if err != nil {
        handler.ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "users":      users,
        "n_users":    len(users),
        "pagination": p,
    })
}

func (h *ManagementHandler) GetUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        handler.ErrorResponse(c, handler.ErrBadID)
        return
    }
    user, err := h.userService.GetUserById(uint(id))
    if err != nil {
        handler.ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}

func (h *ManagementHandler) CreateUser(c *gin.Context) {
    user := &model.User{}
    if err := c.ShouldBindJSON(&user); err != nil {
        handler.ErrorResponse(c, handler.ErrBadRequest)
        return
    }

    user, err := h.userService.CreateUser(user)
    if err != nil {
        handler.ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "user":    user,
        "message": "创建成功",
    })
}

func (h *ManagementHandler) UpdateUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    params := new(model.UpdateUserRequest)
    if err := c.ShouldBindJSON(params); err != nil {
        handler.ErrorResponse(c, err)
        return
    }
    user, err := h.userService.UpdateUser(uint(id), params)
    if err != nil {
        handler.ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}
func (h *ManagementHandler) DeleteUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        handler.ErrorResponse(c, handler.ErrBadID)
        return
    }
    err = h.userService.DeleteUser(uint(id))
    if err != nil {
        handler.ErrorResponse(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "删除成功",
    })
}
