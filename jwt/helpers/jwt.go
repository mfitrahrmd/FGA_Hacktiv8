package helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

func init() {
	os.Setenv("JWT_KEY", "secret")
}

func GenerateToken(id uint, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		log.Fatalln(err.Error())

		return ""
	}

	return tokenString
}
