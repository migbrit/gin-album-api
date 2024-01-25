package http

import (
	"gin-simple-api/infrastructure/http/route"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(ginEngine *gin.Engine) {
	//url groups
	authRouter := ginEngine.Group("/auth")
	albumRouter := ginEngine.Group("/album")

	//handlers
	route.AuthRequestHandler(authRouter)
	route.AlbumRequestHandler(albumRouter)
}
