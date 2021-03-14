package routers

import (
	"github.com/gofiber/fiber/v2"
	"taklif/handlers"
)

// AuthRouter is for routing auth handler
func ApplicationRouter(app *fiber.App) error {

	api := app.Group("/app")

	api.Get("/", handlers.InsertApplicationHandler)

	return nil
}