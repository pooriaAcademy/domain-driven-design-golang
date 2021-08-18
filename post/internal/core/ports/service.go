package ports

import "github.com/pooriaAcademy/event-driven-design-golang/post/internal/core/domain"

type PostService interface {
	CreatePost(Description string) (* domain.Post, error)
	ValidatePostId(PostId string) error
}






