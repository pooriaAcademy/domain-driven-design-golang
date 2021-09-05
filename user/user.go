package user

import (
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/core/service"
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/handlers"
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/repositories"
)

var userRepository = repositories.NewUserMemoryRepository()
var userServiceInstance = &service.UserService{&userRepository}
var HttpHandler = handlers.HttpHandler{userServiceInstance}



