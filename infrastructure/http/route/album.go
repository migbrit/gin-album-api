package route

import (
	"gin-simple-api/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
)

func AlbumRequestHandler(router *gin.RouterGroup) {
	router.GET("/getAll", handler.GetAlbums)
	router.GET("/getById", handler.GetAlbumByID)
	router.POST("/create", handler.PostAlbums)
}
