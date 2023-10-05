// utils/validator.go

package utils

import (
	"regexp"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

// CustomValidator holds a validator instance with custom validation functions
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new custom validator
func NewCustomValidator() *CustomValidator {
	v := validator.New()

	// Register custom email and password validators
	v.RegisterValidation("Email", validateEmail)
	v.RegisterValidation("Password", validatePassword)

	return &CustomValidator{validator: v}
}

// Validate validates a struct based on its tags and custom validation functions
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// validateEmail is a custom email validation function
func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	email = strings.TrimSpace(email)

	// Implement your custom email validation logic here
	// For example, you can use a regular expression to validate the email format
	// Replace the regex pattern with your own pattern
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

// validatePassword is a custom password validation function
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Implement your custom password validation logic here
	// For example, you can check if the password meets certain criteria (e.g., length)
	return len(password) >= 6 // Change this condition as needed
}
