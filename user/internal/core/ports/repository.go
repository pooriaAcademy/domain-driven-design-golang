package ports

import (
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/core/domain"
)

type UserRepository interface {
	SaveUser(user * domain.User) (* domain.User, error)
	GetUserById(UserId string) (* domain.User, error)
	GetUserByUserName(UserId string) (* domain.User, error)
}


