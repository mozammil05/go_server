// utils/password.go

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password and returns the hash
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
	