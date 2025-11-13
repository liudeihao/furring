package main

import (
	db2 "github.com/liudeihao/furring/internal/db"
	"github.com/liudeihao/furring/internal/router"
	"github.com/liudeihao/furring/pkg/log"
)

func main() {
	db, err := db2.NewDB("test.db")
	if err != nil {
		panic("数据库启动失败")
	}
	r := router.NewRouter(db)
	err = r.Run(":8080")
	if err != nil {
		log.Error("程序运行结束")
	}
}
