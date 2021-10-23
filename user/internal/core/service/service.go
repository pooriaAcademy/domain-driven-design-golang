package service

import (
	"errors"
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/core/domain"
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/core/ports"
)

type UserService struct {
	ports.UserRepository
}

func (u UserService) NewUser(UserName string) (*domain.User, error){
	user, err := domain.NewUser(UserName, map[string]bool{}, map[string]bool{})
	if err != nil {
		return nil, err
	}

	user, err = u.UserRepository.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) ToggleFollowingUser(UserId string, FollowingId string) (*domain.User, error){
	user1, err := u.UserRepository.GetUserById(UserId)
	if err != nil {
		return nil, err
	}

	if FollowingId == user1.UserId {
		return nil, errors.New("you can not follow your self")
	}

	user2, err := u.UserRepository.GetUserById(FollowingId)
	if err != nil {
		return nil, err
	}

	if _, ok := user1.FollowingUserIds[user2.UserId]; !ok {
		user1.FollowingUserIds[user2.UserId] = false
	}

	if _, ok := user2.FollowerUserIds[user1.UserId]; !ok {
		user2.FollowerUserIds[user1.UserId] = false
	}
	user1.FollowingUserIds[user2.UserId] = !user1.FollowingUserIds[user2.UserId]
	user2.FollowerUserIds[user1.UserId]  = !user2.FollowerUserIds[user1.UserId]

	user1, err = u.UserRepository.SaveUser(user1)

	if err != nil  {
		return nil, err
	}

	user2, err = u.UserRepository.SaveUser(user2)
	if err != nil  {
		return nil, err
	}

	return user1, nil
}

func (u UserService) GetUser(UserId string) (* domain.User, error){
	return u.UserRepository.GetUserById(UserId)
}




