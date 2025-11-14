package service

import (
    "github.com/liudeihao/furring/internal/model"
    "github.com/liudeihao/furring/pkg/errors"
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userService *UserService
    jwtService  *JWTService
}

func NewAuthService(userService *UserService, jwtService *JWTService) *AuthService {
    return &AuthService{userService: userService, jwtService: jwtService}
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.User, string, error) {
    username := req.Username
    var user *model.User
    var err error
    if username != nil {
        user, err = s.userService.GetUserByUsername(*username)
        if err != nil {
            return nil, "", err
        }
    }
    if req.Email != nil {
        user, err = s.userService.GetUserByEmail(*req.Email)
        if err != nil {
            return nil, "", err
        }
    }
    if user == nil {
        return nil, "", errors.New("必须提供用户名或邮箱才能登录")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return nil, "", errors.New("密码错误")
    }
    token, err := s.jwtService.GenerateToken(user.ID)
    if err != nil {
        return nil, "", err
    }
    return user, token, nil
}

func (s *AuthService) GenerateToken(user *model.User) (string, error) {
    return s.jwtService.GenerateToken(user.ID)
}

func (s *AuthService) Register(req *model.RegisterRequest) (*model.User, error) {
    exists, err := s.userService.usernameExists(req.Email)
    if err != nil {
        return nil, errors.Internal(err, "注册时在数据库中查询邮箱出错")
    }
    if exists {
        return nil, ErrEmailDuplicate
    }
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, errors.Internal(err, "注册，密码加密出错")
    }
    user := &model.User{
        Email:    req.Email,
        Password: string(hashedPassword),
    }
    return user, nil
}
