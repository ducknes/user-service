package handlers

import (
	"github.com/gofiber/fiber/v2"
	"user-service/domain"
	"user-service/service"
	"user-service/tools/usercontext"
)

func RegisterUserHandler(userService service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userCtx := usercontext.New()

		var newUser domain.User
		if err := c.BodyParser(&newUser); err != nil {
			return err
		}

		registeredUser, err := userService.RegisterUser(userCtx, newUser)
		if err != nil {
			return err
		}

		return c.JSON(registeredUser)
	}
}
