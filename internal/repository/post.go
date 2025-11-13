package repository

import (
	"github.com/liudeihao/furring/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Get(id uint) (*model.Post, error)
	Create(post *model.Post) error
	Update(post *model.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Get(id uint) (*model.Post, error) {
	var post model.Post
	if err := r.db.Find(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(model.Post{}, id).Error
}
