package router

import (
    "github.com/gin-gonic/gin"
    "github.com/liudeihao/furring/internal/handler"
    "github.com/liudeihao/furring/internal/handler/admin"
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
    jwtService := service.NewJWTService("this_is_a_secrete_key_that_should_be_put_in_the_environment", "刘德昊")
    authService := service.NewAuthService(userService, jwtService)
    api := r.Group("/api")

    handlers := []handler.Handler{
        admin.NewHandler(userService, postService),
        handler.NewUserHandler(userService, postService),
        handler.NewPostHandler(postService),
        handler.NewAuthHandler(authService),
    }

    for _, h := range handlers {
        h.RegisterHandlers(api)
    }

    return r
}
