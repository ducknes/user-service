package handlers

import (
	"github.com/gofiber/fiber/v2"
	"user-service/service"
	"user-service/tools/usercontext"
)

func GetUserByUsername(userService service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userCtx := usercontext.New()

		user, err := userService.GetUserByUsername(userCtx, c.Query("username"))
		if err != nil {
			return err
		}

		return c.JSON(user)
	}
}
