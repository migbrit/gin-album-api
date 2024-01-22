package http

import (
	"gin-simple-api/infrastructure/http/album"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/all-albums", album.GetAlbums)
	router.GET("/albums", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)
}
