package http

import "github.com/gin-gonic/gin"

func SetupWebServer(router *gin.Engine) error {
	err := router.Run("localhost:8080")
	return err
}
