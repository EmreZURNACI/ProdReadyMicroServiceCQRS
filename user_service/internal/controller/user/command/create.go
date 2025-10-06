package command

import (
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/user/command"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func (h *UserHandler) Create(c *fiber.Ctx) error {

	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := validate.Struct(&user); err != nil {
		return err
	}

	err := h.Service.CreateUser(c.UserContext(), &command.CreateUserRequest{
		User: user,
	})
	if err != nil {
		return err
	}

	return c.Status(200).JSON(map[string]string{
		"status": "OK",
	})

}
