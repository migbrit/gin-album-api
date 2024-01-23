package route

import (
	"gin-simple-api/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
)

func AlbumRequestHandler(router *gin.Engine) {
	router.GET("/all-albums", handler.GetAlbums)
	router.GET("/albums", handler.GetAlbumByID)
	router.POST("/albums", handler.PostAlbums)
}
