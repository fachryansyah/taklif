package models

import (
    "database/sql"
	"fmt"
	"log"
	"os"
	"taklif/utils"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	dbPath := "/etc/data-taklif.db"

	if os.Getenv("IS_DEBUG") == "true" {
		pwd, err := os.Getwd()
		utils.HandleError(err)
		dbPath = fmt.Sprintf("%s/%s", pwd, "databases/data-taklif.db")
	}

	log.Println(dbPath)

    database, _ := sql.Open("sqlite3", dbPath)
	return database
}