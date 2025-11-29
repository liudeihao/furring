package repo

import (
    "github.com/liudeihao/furring/model"
    "gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func (r *userRepository) GetList(opts ...QueryOption) ([]model.User, error) {
    var users []model.User
    result := r.db.Scopes(opts...).Find(&users)
    return users, result.Error
}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
    var user model.User
    result := r.db.First(&user, id)
    return &user, result.Error
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
    var user model.User
    result := r.db.Where("email = ?", email).First(&user)
    return &user, result.Error
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
    var user model.User
    result := r.db.Where("username = ?", username).First(&user)
    return &user, result.Error
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
    result := r.db.Create(user)
    return user, result.Error
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
    result := r.db.Save(user)
    return user, result.Error
}

func (r *userRepository) Delete(user *model.User) error {
    result := r.db.Delete(user)
    return result.Error
}
