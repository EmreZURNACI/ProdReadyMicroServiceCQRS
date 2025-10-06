package command

import "github.com/EmreZURNACI/ProdFullReadyApp_User/internal/service/user/command"

type UserHandler struct {
	Service *command.UserService
}

func NewUserHandler(service *command.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
