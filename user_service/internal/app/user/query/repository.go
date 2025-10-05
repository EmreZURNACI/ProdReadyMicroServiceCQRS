package query

import (
	"context"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
}
