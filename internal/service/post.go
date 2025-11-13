package service

import "github.com/liudeihao/furring/internal/repository"

type PostService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}
