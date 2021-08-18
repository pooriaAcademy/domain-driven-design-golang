package domain

import (
	"github.com/pooriaAcademy/event-driven-design-golang/post"
	"github.com/pooriaAcademy/event-driven-design-golang/user"
)

type Tweet struct {
	TweetId 		string
	PostId 			string
	UserId 			string
	RetweetUserIds 	map[string]bool // map[int]bool acts as a set of people who retweeted the post
	LikedByUserId  	map[string]bool // map[int]bool acts as a set of people who liked the post
}


func NewTweet(PostId string, UserId string) (*Tweet, error){
	err := user.UserServiceInstance.ValidateUserId(UserId)
	if err != nil{
		return nil, err
	}
	err = post.PostService.ValidatePostId(PostId)
	if err != nil {
		return nil, err
	}
	return &Tweet{PostId: PostId, UserId: UserId, RetweetUserIds: map[string]bool{}, LikedByUserId: map[string]bool{}}, nil
}




