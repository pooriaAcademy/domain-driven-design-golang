package service

import "errors"

type UserService struct {

}


func (u UserService) ValidateUserId(UserId string) error {
	return errors.New("user id not valid")
}






