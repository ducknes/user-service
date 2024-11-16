package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"user-service/domain"
	"user-service/service"
	"user-service/tools/usercontext"
)

func DeleteUserHandler(userService service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userCtx := usercontext.New()

		err := userService.DeleteUser(userCtx, c.Query("id"))
		if err != nil && errors.Is(err, domain.NoDocumentAffected) {
			c.Status(fiber.StatusNotFound)
			return nil

		}

		return err
	}
}
