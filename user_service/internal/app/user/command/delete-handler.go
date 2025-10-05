package command

import "context"

type DeleteUserRequest struct {
	ID string `json:"id"`
}
type DeleteUserResponse struct {
}
type DeleteUserHandler struct {
	repository Repository
}

func NewDeleteUserHandler(repo Repository) *DeleteUserHandler {
	return &DeleteUserHandler{
		repository: repo,
	}
}

func (h *DeleteUserHandler) Handle(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {

	if err := h.repository.DeleteUser(ctx, req.ID); err != nil {
		return nil, err
	}
	return &DeleteUserResponse{}, nil
}
