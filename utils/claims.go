// my-auth-app/utils/claims.go
package utils

import "github.com/dgrijalva/jwt-go"

// Claims represents the claims contained within a JWT token
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
