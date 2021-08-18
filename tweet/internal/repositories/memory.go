package repositories

import (
	"errors"
	"github.com/pooriaAcademy/event-driven-design-golang/tweet/internal/core/domain"
)

type MemoryRepository struct {
	Tweets []domain.Tweet
}


func (this * MemoryRepository) SaveTweet(tweet * domain.Tweet) error{
	if tweet.TweetId == "" {
		tweet.TweetId = "NewId"
	}
	this.Tweets = append(this.Tweets, *tweet)
	return nil
}
func (this * MemoryRepository) GetTweet(TweetId string) (*domain.Tweet, error){
	for i := 0 ; i < len(this.Tweets) ; i ++ {
		if this.Tweets[i].TweetId == TweetId{
			return &this.Tweets[i], nil
		}
	}
	return nil, errors.New("tweet doesn't exist")
}





