package query

import (
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/user/query"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *UserHandler) GetID(c *fiber.Ctx) error {

	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(400).JSON(map[string]string{
			"error": "Geçersiz uuid",
		})
	}

	user, err := h.Service.GetUser(c.UserContext(), &query.GetUserRequest{})
	if err != nil {
		return c.Status(400).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]domain.User{
		"user": *user,
	})

}
