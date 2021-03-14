package handlers

import (
	"taklif/providers"
	"taklif/services"
	"taklif/types"
	utils "taklif/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

// LoginHandler is for handle user login
func LoginHandler(c *fiber.Ctx) (error) {

	loginInput := new(types.AuthInput)
	if err := c.BodyParser(loginInput); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	payload := types.AuthInput{
		Email:    loginInput.Email,
		Password: loginInput.Password,
	}

	validate, isValid := services.IsLoginInputValid(payload)
	if !isValid {
		return utils.ResponseValidationError(c, validate, "Validation Error")
	}

	token, err := services.LoginService(payload)
	if err != nil {
		log.Println(err)
		return utils.ResponseBadRequest(c, err.Error(), "Something went wrong")
	}

	return utils.ResponseSuccess(c, token, "Login sukses!")
}

func RegisterHandler(c *fiber.Ctx) error {
	registerInput := new(types.AuthInput)
	if err := c.BodyParser(registerInput); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	payload := types.AuthInput{
		Name:            registerInput.Name,
		Email:           registerInput.Email,
		Password:        registerInput.Password,
		ConfirmPassword: registerInput.ConfirmPassword,
	}

	validate, isValid := services.IsRegisterInputValid(payload)
	if !isValid {
		return utils.ResponseValidationError(c, validate, "Validation Error")
	}

	id, err := services.RegisterService(payload)
	if err != nil {
		log.Println(err)
		return utils.ResponseValidationError(c, "Error", "Validation Error")
	}

	type response struct {
		ID *string
	}

	return utils.ResponseSuccess(c, response{ID: id}, "Registration Sucess!")
}

func CheckTokenHandler(c *fiber.Ctx) error {
	user, err := providers.GetUser(c)
	if err != nil {
		log.Println(err)
		return utils.ResponseError(c, "Token invalid", "Something went wrong")
	}

	return utils.ResponseSuccess(c, &types.LoginIdentity{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, "Token valid!")
}