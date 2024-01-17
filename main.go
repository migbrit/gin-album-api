package main

import (
	"gin-simple-api/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/all-albums", album.GetAlbums)
	router.GET("/albums", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)

	router.Run("localhost:8080")
}
