package command

import (
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/user/command"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *UserHandler) Update(c *fiber.Ctx) error {

	var user domain.User
	if err := c.ParamsParser(&user); err != nil {
		return err
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if _, err := uuid.Parse(user.ID); err != nil {
		return err
	}

	if err := s.Service.UpdateUser(c.UserContext(), &command.UpdateUserRequest{
		User: user,
	}); err != nil {
		return err
	}

	return c.Status(200).JSON(map[string]string{
		"status": "OK",
	})

}
