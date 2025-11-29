package service

import "errors"

var (
    ErrNotAuthorized = errors.New("用户无权限")

    ErrUserNotFound      = errors.New("用户不存在")
    ErrUsernameDuplicate = errors.New("用户名已存在")
    ErrEmailDuplicate    = errors.New("邮箱已被注册")
    ErrPasswordWrong     = errors.New("密码错误")

    ErrPostNotFound = errors.New("帖子不存在")

    ErrCommentNotFound = errors.New("评论不存在")
)
