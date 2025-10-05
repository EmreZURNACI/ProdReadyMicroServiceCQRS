package middleware

import (
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/app/auth/query"
	"github.com/gofiber/fiber/v2"
)

var IsAuthorized = func(c *fiber.Ctx) error {

	var headers map[string][]string = c.GetReqHeaders()
	auth_headers, ok := headers["Authorization"]
	if !ok {
		return c.Status(401).JSON(map[string]string{
			"status":  "Unauthorized",
			"message": "Giriş yapmalısınız",
			"link":    "http://localhost:8081/api/v1/auth/signin",
		})
	}
	var jsonstring string = auth_headers[0]

	_, err := query.NewValidateAccessTokenHandler().Handle(c.UserContext(), &query.ValidateAccessTokenRequest{
		Token: jsonstring,
	})

	if err != nil {
		return c.Status(401).JSON(map[string]string{
			"status":  "Unauthorized",
			"message": "Giriş yapmalısınız",
			"link":    "http://localhost:8081/api/v1/auth/signin",
		})
	}

	return c.Next()
}
