package routers

import (
	"taklif/handlers"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is for routing auth handler
func AuthRouter(app *fiber.App) error {

	api := app.Group("/auth")

	api.Post("/login", handlers.LoginHandler)
	api.Post("/register", handlers.RegisterHandler) // please add middleware

	return nil
}