package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"taklif/services"
	"taklif/types"
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


	id, err := services.RegisterService(types.AuthInput{
		Name: 	  strings.TrimRight(name, "\n"),
		Email:    strings.TrimRight(email, "\n"),
		Password: strings.TrimRight(password, "\n"),
	})

	fmt.Printf("Your user id : %q \n", &id)
	fmt.Println("Successfully setup your account! now you can login with GUI")
}