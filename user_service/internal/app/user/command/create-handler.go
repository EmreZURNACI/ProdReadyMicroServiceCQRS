package command

import (
	"context"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
)

type CreateUserRequest struct {
	User domain.User `json:"user"`
}
type CreateUserResponse struct {
}
type CreateUserHandler struct {
	repository Repository
}

func NewCreateUserHandler(repo Repository) *CreateUserHandler {
	return &CreateUserHandler{
		repository: repo,
	}
}

func (h *CreateUserHandler) Handle(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {

	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID: uuid.String(),
	}

	if err := deepcopier.Copy(&req.User).To(&user); err != nil {
		return nil, err
	}

	if err := h.repository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &CreateUserResponse{}, nil
}
