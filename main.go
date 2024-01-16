package main

import (
	"gin-simple-api/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", album.GetAlbums)

	router.Run("localhost:8080")
}
