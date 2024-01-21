package main

import (
	"gin-simple-api/infrastructure/database"
	"gin-simple-api/infrastructure/http/album"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("An error has occurred trying to get data about .env file -> ", err)
	}

	db, err := database.SetupDatabase()

	if err != nil {
		log.Panic("An error has occurred trying to connect to the database -> ", err)
	}

	defer db.Close()

	album.InitializeDb(db)

	router := gin.Default()
	router.GET("/all-albums", album.GetAlbums)
	router.GET("/albums", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)

	router.Run("localhost:8080")
}
