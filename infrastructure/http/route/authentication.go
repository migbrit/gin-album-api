package route

import (
	"gin-simple-api/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
)

func AuthRequestHandler(router *gin.Engine) {
	router.POST("/auth", handler.GenerateJwtToken)
}
