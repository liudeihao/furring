package service

import (
    "errors"

    "github.com/liudeihao/furring/model"
    "github.com/liudeihao/furring/repo"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func NewUserService() *UserService {
    return &UserService{}
}

type UserService struct{}

func (s *UserService) GetUserPosts(id uint) (*model.UserPostsResponse, error) {
    posts, err := repo.Post.GetList(repo.FilterByUserid(id))
    if len(posts) == 0 {
        return nil, ErrPostNotFound
    }
    if err != nil {
        return nil, err
    }
    ps := make([]model.PostBrief, len(posts))
    for _, post := range posts {
        ps = append(ps, model.PostBrief{
            ID:    post.ID,
            Title: post.Title,
        })
    }
    return &model.UserPostsResponse{
        Posts: ps,
    }, nil
}

func (s *UserService) GetPublicInfoByID(id uint) (*model.UserPublicInfoResponse, error) {
    user, err := repo.User.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &model.UserPublicInfoResponse{
        ID:       user.ID,
        Username: user.Username,
    }, nil
}
func (s *UserService) GetPrivateInfoByID(userid uint, id uint) (*model.UserPrivateInfoResponse, error) {
    if userid != id {
        return nil, ErrNotAuthorized
    }
    user, err := repo.User.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &model.UserPrivateInfoResponse{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
    }, nil
}

func (s *UserService) Login(r model.LoginRequest) (*model.UserLoginResponse, error) {
    username := r.Username
    user, err := repo.User.GetByUsername(username)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
    if err != nil {
        return nil, ErrPasswordWrong
    }

    token, err := GenerateToken(user.ID)
    if err != nil {
        return nil, err
    }

    return &model.UserLoginResponse{
        ID:    user.ID,
        Token: token,
    }, nil
}

func (s *UserService) Register(r model.RegisterRequest) (*model.UserRegisterResponse, error) {
    _, err := repo.User.GetByUsername(r.Username)
    if err == nil {
        // 如果没报错，说明用户存在，所以不可以注册
        return nil, ErrUsernameDuplicate
    } else if !errors.Is(err, gorm.ErrRecordNotFound) {
        // 如果报错了但不是NotFound，则为内部错误
        return nil, err
    }
    _, err = repo.User.GetByEmail(r.Email)
    if err == nil {
        return nil, ErrEmailDuplicate
    } else if !errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, err
    }
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user := &model.User{
        Username: r.Username,
        Email:    r.Email,
        Password: string(hashedPassword),
    }
    user, err = repo.User.Create(user)
    if err != nil {
        return nil, err
    }
    token, err := GenerateToken(user.ID)
    if err != nil {
        return nil, err
    }
    return &model.UserRegisterResponse{
        ID:    user.ID,
        Token: token,
    }, nil
}
