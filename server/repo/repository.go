package repo

import "gorm.io/gorm"

var (
    User    *userRepository
    Post    *postRepository
    Comment *commentRepository
)

func InitRepository(db *gorm.DB) {
    User = &userRepository{db}
    Post = &postRepository{db}
    Comment = &commentRepository{db: db}
}
