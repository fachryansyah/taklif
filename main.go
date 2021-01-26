package main

import (
	"log"
	"taklif/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())

	routers.AuthRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test api")
	})

	app.Listen(":1337")
}
