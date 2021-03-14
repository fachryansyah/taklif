package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"regexp"
	"taklif/models"
	"taklif/types"
	"time"
)

// LoginService is for handle user login
func LoginService(req types.AuthInput) (*types.Token, error) {

	validate, isValid := IsLoginInputValid(types.AuthInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if !isValid {
		log.Println(validate)
		return nil, errors.New("Validation Error")
	}

	findUserByEmail, err := models.FindUserByEmail(req.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(findUserByEmail)

	err = bcrypt.CompareHashAndPassword([]byte(findUserByEmail.Password), []byte(req.Password))
	if err != nil {
		log.Println(err)
		return nil, errors.New("Wrong password!")
	}

	token, err := generateToken(types.LoginIdentity{
		ID:    findUserByEmail.ID,
		Name:  findUserByEmail.Name,
		Email: findUserByEmail.Email,
	}, "user")

	log.Println(fmt.Sprintf("Login token : %s", token))

	return token, nil
}

func RegisterService(req types.AuthInput) (*string, error) {

	validate, isValid := IsRegisterInputValid(types.AuthInput{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if !isValid {
		log.Println(validate)
		return nil, errors.New("Validation Error")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, err
	}

	id, err := models.InsertUser(models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashPassword),
	})
	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf("Registration id : %q", &id))

	return id, nil
}

func IsLoginInputValid(input types.AuthInput) (*types.AuthInput, bool) {
	var validate types.AuthInput
	isValid := true

	if len(input.Email) < 1 {
		validate.Email = "Email can't be null"
		isValid = false
	}

	if len(input.Email) > 255 {
		validate.Email = "Email maximum 255 character"
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(input.Email) {
		validate.Email = "Email format doesn't valif"
		isValid = false
	}

	if len(input.Password) < 1 {
		validate.Password = "Password can't be lower than 1"
		isValid = false
	}

	return &validate, isValid
}

func IsRegisterInputValid(input types.AuthInput) (*types.AuthInput, bool) {
	var validate types.AuthInput
	isValid := true

	if len(input.Name) < 2 {
		validate.Name = "Name can't be lower than 2 character"
		isValid = false
	}

	if len(input.Name) > 255 {
		validate.Name = "Name to long, maximum 255 character"
	}

	if len(input.Email) < 1 {
		validate.Email = "Email can't be null"
		isValid = false
	}

	if len(input.Name) > 255 {
		validate.Name = "Email to long, maximum 255 charackter"
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(input.Email) {
		validate.Email = "Email format doesn't valif"
		isValid = false
	}

	findUserByEmail, err := models.FindUserByEmail(input.Email)
	if err != nil {
		log.Println(err)
		return &validate, isValid
	}

	if len(findUserByEmail.ID) > 1 {
		validate.Email = "Email already used"
		isValid = false
	}

	if len(input.Password) < 6 {
		validate.Password = "Password can't be lower than 6"
		isValid = false
	}

	if input.Password != input.ConfirmPassword {
		validate.Password = "Password not same"
		isValid = false
	}

	return &validate, isValid
}

func generateToken(identity types.LoginIdentity, role string) (*types.Token, error) {
	secretKeyAt := os.Getenv("JWT_KEY")

	// at is access token
	at := jwt.New(jwt.SigningMethodHS256)
	atClaims := at.Claims.(jwt.MapClaims)
	atClaims["identity"] = types.LoginIdentity{
		ID:    identity.ID,
		Email: identity.Email,
		Name:  identity.Name,
	}

	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	atHashed, err := at.SignedString([]byte(secretKeyAt))
	if err != nil {
		return nil, err
	}

	return &types.Token{
		AccessToken:  atHashed,
	}, nil
}