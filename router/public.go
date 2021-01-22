package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	"github.com/markbates/pkger"
)

// PublicRouter is for routing auth handler
func PublicRouter(app *fiber.App) error {

	app.Use("/", filesystem.New(filesystem.Config{
		Root: pkger.Dir("../client/public"),
	}))

	return nil
}
