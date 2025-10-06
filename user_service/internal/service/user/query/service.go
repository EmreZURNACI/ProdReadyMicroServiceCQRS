package query

import (
	"context"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/user/query"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra/mongo"
)

type UserService struct {
	getUserHandler  *query.GetUserHandler
	getUsersHandler *query.GetUsersHandler
}

func NewUserService(repository *mongo.Repository) *UserService {
	return &UserService{
		getUserHandler:  query.NewGetUserHandler(repository),
		getUsersHandler: query.NewGetUsersHandler(repository),
	}
}

func (us *UserService) GetUser(ctx context.Context, req *query.GetUserRequest) (*domain.User, error) {
	res, err := us.getUserHandler.Handle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &res.User, nil
}
func (us *UserService) GetUsers(ctx context.Context, req *query.GetUsersRequest) ([]domain.User, error) {
	res, err := us.getUsersHandler.Handle(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Users, nil
}
