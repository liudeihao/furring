package service

import "github.com/liudeihao/furring/pkg/errors"

var (
	ErrUserNotFound    = errors.New("用户不存在")
	ErrPostNotFound    = errors.New("帖子不存在")
	ErrCommentNotFound = errors.New("评论不存在")
)
