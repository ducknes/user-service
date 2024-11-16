package handlers

import (
	"github.com/gofiber/fiber/v2"
	"user-service/domain"
	"user-service/service"
	"user-service/tools/usercontext"
)

func SaveUserHandler(userService service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userCtx := usercontext.New()

		var user domain.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		return userService.SaveUser(userCtx, user)
	}
}
