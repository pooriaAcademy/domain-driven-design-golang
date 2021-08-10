package ports

import "github.com/pooriaPoorsarvi/event-driven-design-golang/post/internal/core/domain"

type PostRepository interface {
	SavePost(post * domain.Post) error
}






