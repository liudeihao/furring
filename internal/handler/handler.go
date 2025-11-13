package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liudeihao/furring/internal/model"
	"github.com/liudeihao/furring/internal/service"
	"github.com/liudeihao/furring/pkg/pagination"
)

type Handler struct {
	userService *service.UserService
	postService *service.PostService
}

func NewHandler(userService *service.UserService, postService *service.PostService) *Handler {
	return &Handler{
		userService: userService,
		postService: postService,
	}
}

func (h *Handler) RegisterHandlers(r *gin.Engine) {
	users := r.Group("/users")
	users.GET("", h.GetUsersAll)
	users.POST("/", h.CreateUser)
	users.PATCH("/:id", h.UpdateUser)
	users.GET("/:id", h.GetUser)
	users.DELETE("/:id", h.DeleteUser)

}

func (h *Handler) GetUsersAll(c *gin.Context) {
	p := pagination.New(c)

	users, err := h.userService.GetUserList(p.Offset(), p.Limit())
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users":      users,
		"n_users":    len(users),
		"pagination": p,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(c, ErrBadID)
		return
	}
	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}
	if err := h.userService.CreateUser(&user); err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "创建成功",
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
    req := struct {
        Username string `json:"username"`

    }
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ErrorResponse(c, err)
		return
	}
	err = h.userService.UpdateUser(uint(id), &user)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(c, ErrBadID)
		return
	}
	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
