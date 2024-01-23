package handler

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateJwtToken(ginC *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": "An error has occurred generating JWT token."})
		return
	}

	ginC.IndentedJSON(http.StatusOK, gin.H{"Message": "Jwt token generated successfully.", "Token": tokenString})
	return
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return errors.New("Invalid token")
	}

	return nil
}
