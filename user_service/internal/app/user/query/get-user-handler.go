package query

import (
	"context"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
)

type GetUserRequest struct {
	ID string `json:"id"`
}
type GetUserResponse struct {
	User domain.User `json:"user"`
}
type GetUserHandler struct {
	repository Repository
}

func NewGetUserHandler(repo Repository) *GetUserHandler {
	return &GetUserHandler{
		repository: repo,
	}
}

func (h *GetUserHandler) Handle(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {

	user, err := h.repository.GetUser(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &GetUserResponse{
		User: *user,
	}, nil

}
