package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey []byte

type Claims struct {
	jwt.StandardClaims
	User string // Add a User field to store the username or user ID
}

func InitJWT(secret string) {
	jwtKey = []byte(secret)
}

func CreateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
