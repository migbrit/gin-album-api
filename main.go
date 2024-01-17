package main

import (
	"database/sql"
	"gin-simple-api/album"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("An error has occurred trying to get data about .env file -> ", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal("An error has occurred while trying to connect to the database -> ", err)
	} else {
		log.Println("Successfully connected to the database.")
	}

	defer db.Close()

	album.InitializeDb(db)

	router := gin.Default()
	router.GET("/all-albums", album.GetAlbums)
	router.GET("/albums", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)

	router.Run("localhost:8080")
}
