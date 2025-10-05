package command

import (
	"context"
	"time"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/ulule/deepcopier"
)

type UpdateUserRequest struct {
	User domain.User `json:"user"`
}
type UpdateUserResponse struct {
}
type UpdateUserHandler struct {
	repository Repository
}

func NewUpdateUserHandler(repo Repository) *UpdateUserHandler {
	return &UpdateUserHandler{
		repository: repo,
	}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {

	user := domain.User{
		UpdatedAt: time.Now(),
	}

	if err := deepcopier.Copy(&req.User).To(&user); err != nil {
		return nil, err
	}

	user.ID = req.User.ID

	err := h.repository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &UpdateUserResponse{}, nil

}
