package command

import (
	"context"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/user/command"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra/postgres"
)

type UserService struct {
	createUserHandler *command.CreateUserHandler
	updateUserHandler *command.UpdateUserHandler
	deleteUserHandler *command.DeleteUserHandler
}

func NewUserService(pgHandler *postgres.Repository) *UserService {
	return &UserService{
		createUserHandler: command.NewCreateUserHandler(pgHandler),
		updateUserHandler: command.NewUpdateUserHandler(pgHandler),
		deleteUserHandler: command.NewDeleteUserHandler(pgHandler),
	}
}

func (us *UserService) CreateUser(ctx context.Context, req *command.CreateUserRequest) error {
	_, err := us.createUserHandler.Handle(ctx, req)
	return err
}

func (us *UserService) UpdateUser(ctx context.Context, req *command.UpdateUserRequest) error {
	_, err := us.updateUserHandler.Handle(ctx, req)
	return err
}

func (us *UserService) DeleteUser(ctx context.Context, req *command.DeleteUserRequest) error {
	_, err := us.deleteUserHandler.Handle(ctx, req)
	return err
}
