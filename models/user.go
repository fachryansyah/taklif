package models

import (
	"log"
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:""`
}

func InsertUser(user User) (*string, error) {
	db := InitDB()
	uuidVal := uuid.NewString()
	log.Println(uuidVal)
	statment, err := db.Prepare("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	statment.Exec(uuidVal, user.Name, user.Email, user.Password)

	return &uuidVal, nil
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	db := InitDB()
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?",
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}