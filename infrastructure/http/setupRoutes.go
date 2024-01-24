package http

import (
	"gin-simple-api/infrastructure/http/route"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(ginEngine *gin.Engine) {
	authRouter := ginEngine.Group("/auth")
	albumRouter := ginEngine.Group("/album")

	route.AuthRequestHandler(authRouter)
	route.AlbumRequestHandler(albumRouter)
}
