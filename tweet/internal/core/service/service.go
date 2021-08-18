package service

import (
	"github.com/pooriaAcademy/event-driven-design-golang/tweet/internal/core/domain"
	"github.com/pooriaAcademy/event-driven-design-golang/tweet/internal/core/ports"
	"github.com/pooriaAcademy/event-driven-design-golang/user"
)

type TweetService struct {
	ports.TweetRepository
}


func (this * TweetService) NewTweet(PostId string, UserId string) (*domain.Tweet, error){
	return domain.NewTweet(PostId, UserId)
}
func (this * TweetService) ToggleLike(TweetId string, UserId string) error{
	tweet, err := this.TweetRepository.GetTweet(TweetId)
	if err != nil {
		return err
	}
	err = user.UserServiceInstance.ValidateUserId(UserId)
	if err != nil  {
		return err
	}
	val, exists := tweet.LikedByUserId[UserId]
	if exists {
		tweet.LikedByUserId[UserId] = false
	}else{
		tweet.LikedByUserId[UserId] = !val
	}
	err = this.TweetRepository.SaveTweet(tweet)
	if err != nil {
		return err
	}
	return nil
}
func (this * TweetService) ToggleRetweet(TweetId string, UserId string) error{
	tweet, err := this.TweetRepository.GetTweet(TweetId)
	if err != nil {
		return err
	}
	err = user.UserServiceInstance.ValidateUserId(UserId)
	if err != nil  {
		return err
	}
	val, exists := tweet.RetweetUserIds[UserId]
	if exists {
		tweet.RetweetUserIds[UserId] = false
	}else{
		tweet.RetweetUserIds[UserId] = !val
	}
	err = this.TweetRepository.SaveTweet(tweet)
	if err != nil {
		return err
	}
	return nil
}
