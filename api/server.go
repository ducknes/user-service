package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"user-service/api/handlers"
	"user-service/service"
	"user-service/settings"
)

func NewServer(userService service.User) *fiber.App {
	app := fiber.New(fiber.Config{AppName: settings.AppName()})
	app.Use(cors.New(), healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/health",
		ReadinessEndpoint: "/ready",
	}), recover.New())

	app.Get("/user", handlers.GetUserHandler(userService))
	app.Get("/users", handlers.GetUsersHandler(userService))
	app.Post("/user", handlers.SaveUserHandler(userService))
	app.Delete("/user", handlers.DeleteUserHandler(userService))
	app.Put("/user", handlers.UpdateUserHandler(userService))

	app.Post("/user/registration", handlers.RegisterUserHandler(userService))
	app.Get("/user/by-username", handlers.GetUserByUsername(userService))

	return app
}
