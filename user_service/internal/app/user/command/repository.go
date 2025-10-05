package command

import (
	"context"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
}

//internal içinde dizin adı ile import etmek içingithub.com/EmreZURNACI/ProdFullReadyApp_User
// import edilecek olan paket de
// import edilen yer de  internla içinde olması gerekir.
