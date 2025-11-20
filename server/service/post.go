package service

import (
    "errors"

    "github.com/liudeihao/furring/model"
    "github.com/liudeihao/furring/repo"
    "gorm.io/gorm"
)

type PostService struct{}

func NewPostService() *PostService {
    return &PostService{}
}

func (s *PostService) Create(r *model.PostCreateRequest) (uint, error) {
    post := model.Post{
        UserID:  r.UserID,
        Title:   r.Title,
        Content: r.Content,
    }
    p, err := repo.Post.Create(&post)
    if err != nil {
        return 0, err
    }
    return p.ID, nil
}

func (s *PostService) GetByID(id uint) (*model.PostResponse, error) {
    post, err := repo.Post.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrPostNotFound
        }
        return nil, err
    }
    user, err := repo.User.GetByID(post.UserID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrPostNotFound
        }
        return nil, err
    }
    return &model.PostResponse{
        UserInfo: model.UserInfo{
            UserID:   user.ID,
            Username: user.Username,
        },
        ID:      post.ID,
        Title:   post.Title,
        Content: post.Content,
    }, nil
}

func (s *PostService) Update(userid uint, r model.PostUpdateRequest) error {
    post, err := repo.Post.GetByID(r.PostID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return ErrPostNotFound
        }
        return err
    }
    if post.UserID != userid {
        return ErrNotAuthorized
    }
    if r.Title != nil {
        post.Title = *r.Title
    }
    if r.Content != nil {
        post.Content = *r.Content
    }
    post, err = repo.Post.Update(post)
    return err
}

func (s *PostService) Delete(userid uint, postid uint) error {
    post, err := repo.Post.GetByID(postid)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return ErrPostNotFound
        }
        return err
    }
    if post.UserID != userid {
        return ErrNotAuthorized
    }
    err = repo.Post.Delete(post.ID)
    if err != nil {
        return err
    }
    return nil
}
