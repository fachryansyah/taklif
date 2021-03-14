package services

import (
	"log"
	"taklif/models"
)

func InsertApplicatoinService(req models.Application) (*string, error) {
	id, err := models.InsertApplication(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return id, nil
}
