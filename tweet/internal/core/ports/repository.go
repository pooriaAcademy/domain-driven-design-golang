package ports

import "github.com/pooriaAcademy/event-driven-design-golang/tweet/internal/core/domain"

type PostRepository interface {
	SavePost(post * domain.Tweet) error
}






