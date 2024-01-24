package route

import (
	"gin-simple-api/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
)

func AuthRequestHandler(router *gin.RouterGroup) {

	router.POST("/createToken", handler.GenerateJwtToken)
}
