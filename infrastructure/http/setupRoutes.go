package http

import (
	"gin-simple-api/infrastructure/http/route"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(ginEngine *gin.Engine) {
	route.AuthRequestHandler(ginEngine)
	route.AlbumRequestHandler(ginEngine)
}
