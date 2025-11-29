package repo

import (
    "github.com/liudeihao/furring/model"
    "gorm.io/gorm"
)

type postRepository struct {
    db *gorm.DB
}

func (r *postRepository) GetByID(id uint) (*model.Post, error) {
    var post model.Post
    result := r.db.First(&post, id)
    return &post, result.Error
}

func (r *postRepository) Create(post *model.Post) (*model.Post, error) {
    result := r.db.Create(post)
    return post, result.Error
}

func (r *postRepository) Update(post *model.Post) (*model.Post, error) {
    result := r.db.Save(post)
    return post, result.Error
}

func (r *postRepository) Delete(id uint) error {
    result := r.db.Delete(&model.Post{}, id)
    return result.Error
}

func (r *postRepository) GetList(opts ...QueryOption) ([]model.Post, error) {
    var posts []model.Post
    result := r.db.Scopes(opts...).Find(&model.Post{})
    return posts, result.Error
}
