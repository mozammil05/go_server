package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey []byte

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
