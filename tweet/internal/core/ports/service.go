package ports

import "github.com/pooriaAcademy/event-driven-design-golang/tweet/internal/core/domain"

type PostService interface {
	NewTweet(PostId string, UserId string) (*domain.Tweet, error)
	ToggleLike(TweetId string, UserId string) error
	ToggleRetweet(TweetId string, UserId string) error
}






