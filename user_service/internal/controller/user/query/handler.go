package query

import "github.com/EmreZURNACI/ProdFullReadyApp_User/internal/service/user/query"

type UserHandler struct {
	Service *query.UserService
}

func NewUserHandler(service *query.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
