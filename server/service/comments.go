package service

import (
    "errors"

    "github.com/liudeihao/furring/model"
    "github.com/liudeihao/furring/repo"
    "gorm.io/gorm"
)

type CommentService struct{}

func NewCommentService() *CommentService {
    return &CommentService{}
}

func (*CommentService) GetComment(id uint) (*model.CommentResponse, error) {
    comment, err := repo.Comment.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrCommentNotFound
        }
        return nil, err
    }
    user, err := repo.User.GetByID(comment.UserID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrCommentNotFound
        }
        return nil, err
    }
    return &model.CommentResponse{
        Username: user.Username,
        Content:  comment.Content,
    }, nil
}

func (s *CommentService) Comment(r model.CommentCreateRequest) (uint, error) {
    comment := &model.Comment{
        UserID:  r.UserID,
        PostID:  r.PostID,
        Content: r.Content,
    }

    _, err := repo.Post.GetByID(r.PostID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return 0, ErrPostNotFound
        }
        return 0, err
    }

    c, err := repo.Comment.Create(comment)
    if err != nil {
        return 0, err
    }
    return c.ID, nil
}

func (s *CommentService) Delete(uid, id uint) error {
    comment, err := repo.Comment.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return ErrCommentNotFound
        }
        return err
    }
    if comment.UserID != uid {
        return ErrNotAuthorized
    }
    err = repo.Comment.Delete(id)
    if err != nil {
        return err
    }
    return nil
}
