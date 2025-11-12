package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func NewDB(source string) (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(source), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
    if err != nil {
        return nil, err
    }
    // sqlite 默认不启用外键约束
    db.Exec("PRAGMA foreign_keys = ON;")
    return db, nil
}
