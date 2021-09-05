package utils

import "errors"

type CustomErrors struct {
	Status CustomErrorsStatus
	Err error
}

func NewCustomErrors(stat CustomErrorsStatus) * CustomErrors {
	return &CustomErrors{
		Status: stat,
		Err: errors.New(customErrorsMessage[stat]),
	}
}

func (c CustomErrors) Error() string {
	return c.Err.Error()
}


type CustomErrorsStatus int

const (
	UserNotFoundStatus CustomErrorsStatus  = iota
	ServerErrorStatus
)



var customErrorsMessage = map[CustomErrorsStatus]string{
	UserNotFoundStatus: "no user with the specified features were found",
	ServerErrorStatus: "server error happened, please contact as as soon as possible",
}

func (c CustomErrorsStatus) Error() string{
	return customErrorsMessage[c]
}

















