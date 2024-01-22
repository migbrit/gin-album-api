package main

import (
	"gin-simple-api/infrastructure/database"
	"gin-simple-api/infrastructure/http"
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

	http.SetupRoutes(router)
	if http.SetupWebServer(router); err != nil {
		log.Panic("An error has occurred trying to setup the web server -> ", err)
	}
}
