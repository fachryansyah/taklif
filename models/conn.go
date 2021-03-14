package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func InitDB() *sql.DB {
	dbPath := os.Getenv("HOME") + "/.taklif/data.db"

	//if os.Getenv("IS_DEBUG") == "true" {
	//	pwd, err := os.Getwd()
	//	utils.HandleError(err)
	//	dbPath = fmt.Sprintf("%s/%s", pwd, "databases/data-taklif.db")
	//}

	log.Println(dbPath)

    database, _ := sql.Open("sqlite3", dbPath)
	return database
}