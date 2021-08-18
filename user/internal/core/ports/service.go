package ports



type UserService interface {
	ValidateUserId(UserId string) error
}


