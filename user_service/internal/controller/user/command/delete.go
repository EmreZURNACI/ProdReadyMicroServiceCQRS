package command

import (
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/user/command"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *UserHandler) Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return err
	}

	err := h.Service.DeleteUser(c.UserContext(), &command.DeleteUserRequest{
		ID: id,
	})
	if err != nil {
		return err
	}

	return c.Status(200).JSON(map[string]string{
		"status": "OK",
	})
}
