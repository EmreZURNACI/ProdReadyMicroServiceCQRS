package query

import (
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/user/query"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) Get(c *fiber.Ctx) error {

	users, err := h.Service.GetUsers(c.UserContext(), &query.GetUsersRequest{})
	if err != nil {
		return c.Status(200).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(map[string][]domain.User{
		"users": users,
	})

}
