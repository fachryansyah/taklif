package main

import (
	"fmt"
	"taklif/models"

	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Migating db sample..")

	db := models.InitDB()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var migration []string
	pathMigration := fmt.Sprintf("%s/%s", pwd, "databases/migration")
	err = filepath.Walk(pathMigration, func(path string, info os.FileInfo, err error) error {
		migration = append(migration, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for i, file := range migration {
		if i != 0 && i < len(migration)-1 {
			log.Println(fmt.Sprintf("migrating : %q", file))
			content, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			_, err = db.Exec(string(content))
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		}
	}

	log.Println("Migration complete!")
}
