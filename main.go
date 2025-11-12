package main

import "github.com/liudeihao/furring/log"

func main() {
	db, err := NewDB("test.db")
	if err != nil {
		panic("数据库启动失败")
	}
	r := NewRouter(db)
	err = r.Run(":8080")
	if err != nil {
		log.Error("程序运行结束")
	}
}
