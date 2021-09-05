package ports

import "github.com/pooriaAcademy/event-driven-design-golang/user/internal/core/domain"

type UserService interface {
	NewUser(UserName string) (*domain.User, error)
	ToggleFollowingUser(UserId string, FollowerId string) (*domain.User, error)
	GetUser(UserId string) (* domain.User, error)
}


