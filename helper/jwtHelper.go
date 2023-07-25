package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userid string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["userid"] = userid
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "Signing Error", err
	}
	return tokenString, nil
}
