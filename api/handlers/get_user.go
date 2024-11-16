package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"user-service/domain"
	"user-service/service"
	"user-service/tools/usercontext"
)

func GetUserHandler(userService service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userCtx := usercontext.New()

		user, err := userService.GetUser(userCtx, c.Query("id"))
		if err != nil {
			if errors.Is(err, domain.NoDocuments) {
				c.Status(fiber.StatusNoContent)
				return nil
			}

			return err
		}

		return c.JSON(user)
	}
}
