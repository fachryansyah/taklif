package handler

import (
	"taklif/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginHandler(c *fiber.Ctx) error {
	
	return utils.ResponseSuccess(c, &LoginResponse{}, "Login sukses!")
}
