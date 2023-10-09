// // utils/validator.go

// package utils

// import (
// 	"regexp"
// 	"strings"

// 	"github.com/go-playground/locales/en"
// 	ut "github.com/go-playground/universal-translator"
// 	validator "github.com/go-playground/validator/v10"
// )

// // CustomValidator holds a validator instance with custom validation functions and error messages
// type CustomValidator struct {
// 	validator  *validator.Validate
// 	translator ut.Translator
// }

// // NewCustomValidator creates a new custom validator with custom error messages
// func NewCustomValidator() *CustomValidator {
// 	v := validator.New()

// 	// Register custom email and password validators
// 	v.RegisterValidation("Email", validateEmail)
// 	v.RegisterValidation("Password", validatePassword)

// 	// Create an English translator for error messages
// 	english := en.New()
// 	uni := ut.New(english, english)
// 	trans, _ := uni.GetTranslator("en")

// 	// Register custom error messages for specific fields and tags
// 	v.RegisterTranslation("email_required", trans, func(ut ut.Translator) error {
// 		return ut.Add("email_required", "Email is required.", true)
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("email_required", fe.Field())
// 		return t
// 	})

// 	v.RegisterTranslation("email_format", trans, func(ut ut.Translator) error {
// 		return ut.Add("email_format", "Invalid email format.", true)
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("email_format", fe.Field())
// 		return t
// 	})

// 	// Register more custom error messages as needed...

// 	return &CustomValidator{validator: v, translator: trans}
// }

// // Validate validates a struct based on its tags and custom validation functions
// func (cv *CustomValidator) Validate(i interface{}) error {
// 	return cv.validator.Struct(i)
// }

// // validateEmail is a custom email validation function
// func validateEmail(fl validator.FieldLevel) bool {
// 	email := fl.Field().String()
// 	email = strings.TrimSpace(email)

// 	// Implement your custom email validation logic here
// 	// For example, you can use a regular expression to validate the email format
// 	// Replace the regex pattern with your own pattern
// 	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
// 	match, _ := regexp.MatchString(emailPattern, email)
// 	return match
// }

// // validatePassword is a custom password validation function
// func validatePassword(fl validator.FieldLevel) bool {
// 	password := fl.Field().String()

//		// Implement your custom password validation logic here
//		// For example, you can check if the password meets certain criteria (e.g., length)
//		return len(password) >= 6 // Change this condition as needed
//	}
package utils

import (
	"regexp"
	"strings"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// CustomValidator holds a validator instance with custom validation functions
type CustomValidator struct {
	validator *validator.Validate
}

var translator ut.Translator

// NewCustomValidator creates a new custom validator
func NewCustomValidator() *CustomValidator {
	v := validator.New()
	// Register custom email and password validators
	v.RegisterValidation("email", validateEmail)
	v.RegisterValidation("password", validatePassword)

	// Create a new translator with the desired locale (e.g., en_US)
	en := en_US.New() // Replace with your preferred locale
	uni := ut.New(en, en)
	translator, _ = uni.GetTranslator("en_US") // Replace with your preferred locale

	// Define custom error messages for specific fields and tags
	v.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	v.RegisterTranslation("email", translator, func(ut ut.Translator) error {
		return ut.Add("email", "Invalid email format for {0}.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	// Register more custom error messages as needed...

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
