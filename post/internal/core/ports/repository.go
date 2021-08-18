package ports

import "github.com/pooriaAcademy/event-driven-design-golang/post/internal/core/domain"

type PostRepository interface {
	SavePost(post * domain.Post) error
	NewPostId(post * domain.Post) (string, error)
}






