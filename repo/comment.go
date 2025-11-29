package repo

import (
    "github.com/liudeihao/furring/model"
    "gorm.io/gorm"
)

type commentRepository struct {
    db *gorm.DB
}

func (r *commentRepository) Create(comment *model.Comment) (*model.Comment, error) {
    result := r.db.Create(comment)
    return comment, result.Error
}

func (r *commentRepository) GetByID(id uint) (*model.Comment, error) {
    comment := &model.Comment{}
    result := r.db.Where("id = ?", id).First(comment)
    return comment, result.Error
}

func (r *commentRepository) GetList(opts ...QueryOption) ([]model.Comment, error) {
    var comments []model.Comment
    result := r.db.Scopes(opts...).Find(&comments)
    return comments, result.Error
}

func (r *commentRepository) Delete(id uint) error {
    result := r.db.Delete(&model.Comment{}, id)
    return result.Error
}
