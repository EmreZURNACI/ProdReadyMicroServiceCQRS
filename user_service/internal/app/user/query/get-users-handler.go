package query

import (
	"context"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
)

type GetUsersRequest struct {
}
type GetUsersResponse struct {
	Users []domain.User `json:"users"`
}
type GetUsersHandler struct {
	repository Repository
}

func NewGetUsersHandler(repo Repository) *GetUsersHandler {
	return &GetUsersHandler{
		repository: repo,
	}
}

func (h *GetUsersHandler) Handle(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {

	users, err := h.repository.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return &GetUsersResponse{
		Users: users,
	}, nil

}
