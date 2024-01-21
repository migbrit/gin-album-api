package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func SetupDatabase() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("dbName: ", dbName)

	dbConnectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal("An error has occurred while trying to connect to the database -> ", err)
	} else {
		log.Println("Successfully connected to the database.")
	}

	return db, err
}
