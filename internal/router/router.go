package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liudeihao/furring/internal/handler"
	"github.com/liudeihao/furring/internal/repository"
	"github.com/liudeihao/furring/internal/service"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)
	h := handler.NewHandler(userService, postService)
	h.RegisterHandlers(r)

	return r
}
