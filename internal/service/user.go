package service

import (
    "strings"

    "github.com/liudeihao/furring/internal/model"
    "github.com/liudeihao/furring/internal/repository"
    "github.com/liudeihao/furring/pkg/errors"
    "gorm.io/gorm"
)

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserList(offset, limit int) ([]model.User, error) {
    users, err := s.userRepo.FindMultiple(offset, limit)
    if err != nil {
        return nil, errors.Internal(err, "获取用户列表失败")
    }
    return users, nil
}

func (s *UserService) GetUserById(id uint) (*model.User, error) {
    user, err := s.userRepo.FindByID(id)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    } else if err != nil {
        return nil, errors.Internal(err, "通过ID获取用户失败")
    }
    return user, nil
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
    user, err := s.userRepo.FindByUsername(username)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    } else if err != nil {
        return nil, errors.Internal(err, "通过Username获取用户失败")
    }
    return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
    user, err := s.userRepo.FindByEmail(email)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    } else if err != nil {
        return nil, errors.Internal(err, "通过Email获取用户失败")
    }
    return user, nil
}

func (s *UserService) usernameExists(username string) (bool, error) {
    user, err := s.userRepo.FindByUsername(username)
    if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
        return false, err
    }
    return user != nil, nil
}
func (s *UserService) emailExists(email string) (bool, error) {
    email = strings.ToLower(email)
    user, err := s.userRepo.FindByEmail(email)
    if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
        return false, err
    }
    return user != nil, nil
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
    // 整理输入
    user.Email = strings.ToLower(user.Email)

    // 检查是否已存在
    usernameExist, err := s.usernameExists(user.Username)
    if err != nil {
        return nil, errors.Internal(err, "检测用户名是否存在失败")
    }
    if usernameExist {
        return nil, ErrUsernameDuplicate
    }
    emailExist, err := s.emailExists(user.Email)
    if err != nil {
        return nil, errors.Internal(err, "检测email是否存在失败")
    }
    if emailExist {
        return nil, ErrEmailDuplicate
    }

    // 创建用户
    rowsAffected, err := s.userRepo.Create(user)
    if rowsAffected == 0 || err != nil {
        return user, errors.Internal(err, "创建用户失败")
    }
    return user, nil
}

func (s *UserService) UpdateUser(id uint, params *model.UpdateUserRequest) (*model.User, error) {
    user, err := s.userRepo.FindByID(id)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    } else if err != nil {
        return nil, errors.Internal(err, "更新用户前查询失败")
    }
    if params.Username != nil {
        newname := *params.Username
        exists, err := s.usernameExists(newname)
        if err != nil {
            return nil, errors.Internal(err, "检测用户名是否存在失败")
        }
        if exists && newname != *params.Username {
            return nil, ErrUsernameDuplicate
        }
        user.Username = newname
    }

    rowsAffected, err := s.userRepo.Update(user)
    if err != nil {
        return nil, errors.Internal(err, "更新用户失败")
    }
    if rowsAffected == 0 {
        return nil, ErrUserNotFound
    }
    return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
    rowsAffected, err := s.userRepo.Delete(id)
    if err != nil {
        return errors.Internal(err, "删除用户失败")
    }
    if rowsAffected == 0 {
        return ErrUserNotFound
    }
    return nil
}
