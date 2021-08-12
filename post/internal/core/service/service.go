package service

import (
	"github.com/pooriaAcademy/event-driven-design-golang/post/internal/core/domain"
	"github.com/pooriaAcademy/event-driven-design-golang/post/internal/core/ports"
)

type PostService struct {
	ports.PostRepository
}


func NewPostService(repository ports.PostRepository) * PostService{
	return &PostService{PostRepository: repository}
}


func (this * PostService) CreatePost(Description string) (* domain.Post, error) {
	post, err := domain.NewPost(Description)
	if err != nil{
		return nil, err
	}
	err = this.PostRepository.SavePost(post)
	if err != nil{
		return nil, err
	}
	return post, nil
}



func (this * PostService) ValidatePostId(PostId string) bool{
	return true
}



