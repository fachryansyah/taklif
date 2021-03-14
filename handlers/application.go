package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"taklif/models"
	"taklif/services"
	"taklif/utils"
)

func InsertApplicationHandler (c *fiber.Ctx) error {
	application := new(models.Application)
	if err := c.BodyParser(application); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	payload := models.Application{
		Pid: application.Pid,
		Name: application.Name,
		Runtime: application.Runtime,
		Path: application.Path,
	}

	id, err := services.InsertApplicatoinService(payload)
	if err != nil {
		log.Println(err)
		return utils.ResponseError(c, err.Error(), "Something went wrong")
	}

	return utils.ResponseSuccess(c, id, "Successfully!")
}