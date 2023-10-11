package utils

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Email       string `json:"email"`
	Role        string `json:"role"`
	CustomClaim string `json:"custom_claim"`
	jwt.StandardClaims
}
