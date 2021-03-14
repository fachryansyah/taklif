package models

import (
	"github.com/google/uuid"
	"log"
	"time"
)

type Application struct {
	ID        string `json:"id"`
	Pid       string `json:"pid"`
	Name      string `json:"name"`
	Runtime   string `json:"runtime"` // 0 = go, 1 = node, 2 = php, 3 = python
	Path      string `json:"path"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"created_at"`
}

func GetApplications() ([]*Application, error) {
	var applications []*Application

	db := InitDB()
	query := "SELECT id, pid, name, runtime, path, created_at FROM applications"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var item Application
		var created_at time.Time
		err = rows.Scan(&item.ID, &item.Pid, &item.Name, &item.Runtime, &item.Path, &created_at)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		item.CreatedAt = created_at.Unix()

		applications = append(applications, &item)
	}

	return applications, nil
}

func FindApplicationById(id string) (*Application, error) {
	var application Application
	var created_at time.Time
	db := InitDB()
	err := db.QueryRow(`
		SELECT id, pid, name, runtime, path, created_at FROM applications WHERE id = ?
	`, id,
	).Scan(&application.ID, &application.Pid, &application.Name, &application.Runtime, &application.Path, &created_at)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	application.CreatedAt = created_at.Unix()

	return &application, nil
}

func InsertApplication(application Application) (*string, error) {
	db := InitDB()
	uuidVal := uuid.NewString()
	statment, err := db.Prepare(`
		INSERT INTO applications
		(
			id,
		    pid,
		    name,
		    runtime,
		 	path
		) VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	statment.Exec(uuidVal, application.Pid, application.Name, application.Runtime, application.Path)

	return &uuidVal, nil
}

func UpdateApplication(application Application) error {
	return nil
}

func DeleteApplication(id string) error {
	return nil
}
