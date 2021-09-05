package repositories

import (
	"errors"
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/core/domain"
	"strconv"
)

type UserMemoryRepository struct {
	Users []domain.User
}

func NewUserMemoryRepository () UserMemoryRepository {
	return UserMemoryRepository{Users: []domain.User{}}
}



func(t * UserMemoryRepository) GetUserById(UserId string) (* domain.User, error){
	for i := 0; i < len(t.Users); i++ {
		if t.Users[i].UserId == UserId {
			return &t.Users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func(t * UserMemoryRepository) GetUserByUserName(UserName string) (* domain.User, error){
	for i := 0; i < len(t.Users); i++ {
		if t.Users[i].UserName == UserName {
			return &t.Users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func(t * UserMemoryRepository) SaveUser(user * domain.User) (* domain.User, error){
	i := 0
	for ; i < len(t.Users); i++ {
		if t.Users[i].UserId == user.UserId || t.Users[i].UserName == user.UserName{
			user.UserId = t.Users[i].UserId
			t.Users[i] = *user
			return &t.Users[i], nil
		}
	}
	user.UserId = strconv.Itoa(len(t.Users))
	t.Users = append(t.Users, *user)
	return &t.Users[i], nil
}









