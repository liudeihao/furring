package model

import (
    "time"
)

type User struct {
    Model
    Username string `gorm:"size:32;unique" json:"username"`
    Email    string `gorm:"size:255;unique" json:"email"`
    Password string `gorm:"size:255" json:"-"`
}

type UserToken struct {
    Model
    UserID   uint      `gorm:"not null" json:"user_id"`
    Token    string    `gorm:"size:255" json:"token"`
    ExpireAt time.Time `gorm:"not null" json:"expire_at"`
}

type Post struct {
    Model
    UserID  uint   `gorm:"not null" json:"user_id"`
    Title   string `gorm:"size:255" json:"title"`
    Content string `gorm:"type:text" json:"content"`
}

type Comment struct {
    Model
    UserID  uint   `gorm:"not null" json:"user_id"`
    PostID  uint   `gorm:"not null" json:"post_id"`
    Content string `gorm:"type:text" json:"content"`
}
