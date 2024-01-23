package middleware

import (
	"gin-simple-api/infrastructure/http/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware(ginC *gin.Context) {
	tokenString := ginC.GetHeader("Authorization")

	err := handler.VerifyToken(tokenString)
	if err != nil {
		ginC.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ginC.Abort()
		return
	}

	ginC.Next()
}
