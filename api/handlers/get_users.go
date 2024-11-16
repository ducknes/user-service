package handlers

import (
	"errors"
	"strconv"
	"user-service/domain"
	"user-service/service"
	"user-service/tools/usercontext"

	"github.com/gofiber/fiber/v2"
)

func GetUsersHandler(userService service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userCtx := usercontext.New()

		limit, cursor := parseParams(c)

		users, err := userService.GetUsers(userCtx, limit, cursor)
		if err != nil {
			if errors.Is(err, domain.NoDocuments) {
				c.Status(fiber.StatusNotFound)
				return nil
			}

			return err
		}

		return c.JSON(users)
	}
}

func parseParams(c *fiber.Ctx) (int64, string) {
	limit := c.Query("limit")
	cursor := c.Query("cursor")

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt == 0 {
		limitInt = 10
	}

	return int64(limitInt), cursor
}
