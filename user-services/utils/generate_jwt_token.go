package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a JWT token for a user
func GenerateJWT(username, role string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	secretKey := os.Getenv("SECRET_KEY")

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
