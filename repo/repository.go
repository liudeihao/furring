package repo

import "gorm.io/gorm"

var (
    User    *userRepository
    Post    *postRepository
    Comment *commentRepository
    DB      *gorm.DB
)

func InitRepository(db *gorm.DB) {
    DB = db
    User = &userRepository{db}
    Post = &postRepository{db}
    Comment = &commentRepository{db: db}
}
