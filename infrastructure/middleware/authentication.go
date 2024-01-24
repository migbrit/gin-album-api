package middleware

import (
	"gin-simple-api/infrastructure/http/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware(ginC *gin.Context) {
	path := ginC.Request.URL.Path

	if strings.Contains(path, "/createToken") {
		ginC.Next()
		return
	}

	tokenString := ginC.GetHeader("Authorization")

	err := handler.VerifyToken(tokenString)
	if err != nil {
		ginC.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ginC.Abort()
		return
	}

	ginC.Next()
}
