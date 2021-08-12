package domain

import "github.com/pooriaAcademy/event-driven-design-golang/user"

type Tweet struct {
	PostId 			string
	UserId 			string
	RetweetUserIds 	[]string
	LikedByUserId  	[]string
}


func NewTweet(PostId string, UserId string) (*Tweet, error){
	user.UserServiceInstance.ValidateUserId()
	return nil, nil
}




