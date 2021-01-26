package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"taklif/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nWelcome to TAKLIF, Setup your first account for login on Dashboard GUI!\n")

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your mail: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
	}

	_, err = models.InsertUser(models.User{
		Name:     strings.TrimRight(name, "\n"),
		Email:    strings.TrimRight(email, "\n"),
		Password: string(hashPassword),
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Print("Successfully setup your account")
}