package ports

import "github.com/pooriaAcademy/event-driven-design-golang/tweet/internal/core/domain"

type TweetRepository interface {
	SaveTweet(tweet * domain.Tweet) error
	GetTweet(TweetId string) (*domain.Tweet, error)
}






