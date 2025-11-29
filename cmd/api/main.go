package main

import (
    "context"
    "errors"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/liudeihao/furring/config"
    "github.com/liudeihao/furring/database"
    "github.com/liudeihao/furring/handler"
    "github.com/liudeihao/furring/initialize"
    "github.com/liudeihao/furring/repo"
    "github.com/liudeihao/furring/router"
    "github.com/liudeihao/furring/service"
)

func main() {
    initialize.LoadConfig()
    c := config.Instance

    db := database.New(c.DB.Driver, c.DB.DSN)
    repo.InitRepository(db)

    userService := service.NewUserService()
    postService := service.NewPostService()
    commentService := service.NewCommentService()

    handlers := []handler.Handler{
        handler.NewUserHandler(userService),
        handler.NewPostHandler(postService),
        handler.NewCommentHandler(commentService),
    }

    r := router.New(handlers)

    srv := &http.Server{
        Handler: r.Handler(),
        Addr:    fmt.Sprintf("%s:%s", config.Instance.Server.Addr, config.Instance.Server.Port),
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Fatalf("Listen: %s\n", err)
        }
    }()

    // 优雅退出
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    s := <-quit
    log.Println("收到操作系统信号：" + s.String())

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err := srv.Shutdown(ctx)
    if err != nil {
        log.Println("服务器关闭时遇到错误：", err)
    }
    log.Print("服务器成功关闭")
}
