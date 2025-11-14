package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email    string `json:"email" gorm:"uniqueIndex"`
    Password string `json:"-"`
    Username string `json:"username" gorm:"uniqueIndex"`
}

type Post struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
    User    *User  `json:"user,omitempty"`
}

type Comment struct {
    gorm.Model
    PostID uint  `json:"post_id"`
    Post   *Post `json:"post,omitempty"`
    UserID uint  `json:"user_id"`
    User   *User `json:"user,omitempty"`
}
