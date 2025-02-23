package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "aP9!x@2#b$5^d&7*e(0)f+3=g~4_h-6"

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
