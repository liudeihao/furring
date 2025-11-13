package repository

import (
	"github.com/liudeihao/furring/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetMultiple(offset, limit int) ([]model.User, error)
	Create(user *model.User) (int64, error)
	Update(user *model.User) (int64, error)
	Delete(id uint) (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetMultiple(offset, limit int) ([]model.User, error) {
	var users []model.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Get(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) (int64, error) {
	result := r.db.Create(user)
	return result.RowsAffected, result.Error
}

func (r *userRepository) Update(user *model.User) (int64, error) {
	result := r.db.Save(user)
	return result.RowsAffected, result.Error
}

func (r *userRepository) Delete(id uint) (int64, error) {
	result := r.db.Delete(&model.User{}, id)
	return result.RowsAffected, result.Error
}
