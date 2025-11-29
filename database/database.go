package database

import (
    "github.com/liudeihao/furring/model"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func New(driver, addr string) *gorm.DB {
    switch driver {
    case "sqlite":
        db, err := gorm.Open(sqlite.Open(addr), &gorm.Config{})
        if err != nil {
            panic("连接数据库失败：" + err.Error())
        }
        err = db.AutoMigrate(
            &model.User{},
            &model.Post{},
            &model.Comment{},
        )
        if err != nil {
            panic("数据库迁移失败：" + err.Error())
        }
        return db
    default:
        panic("尚未支持其他数据库")
    }
}
