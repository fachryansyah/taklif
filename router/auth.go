package router

import (
	"taklif/handler"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is for routing auth handler
func AuthRouter(app *fiber.App) error {

	api := app.Group("/auth")

	api.Get("/login", handler.LoginHandler)

	return nil
}
