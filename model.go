package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
	Username string `json:"username"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
	User    *User  `json:"user,omitempty"`
}

type Comment struct {
	gorm.Model
	PostID int   `json:"post_id"`
	Post   *Post `json:"post,omitempty"`
	UserID int   `json:"user_id"`
	User   *User `json:"user,omitempty"`
}
