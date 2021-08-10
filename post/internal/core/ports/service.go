package ports

import "github.com/pooriaPoorsarvi/event-driven-design-golang/post/internal/core/domain"

type PostService interface {
	CreatePost(Description string) (* domain.Post, error)
	ValidatePostId(PostId string) bool
}






