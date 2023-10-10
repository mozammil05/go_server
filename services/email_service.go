// services/email_service.go

package services

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendResetEmail sends a password reset email to the specified recipient
func SendResetEmail(recipientEmail, resetLink string) error {
	// Load SendGrid API key from environment variable
	sendGridAPIKey := os.Getenv("SENDGRID_API_KEY")
	if sendGridAPIKey == "" {
		return fmt.Errorf("SendGrid API key not found in environment variable SENDGRID_API_KEY")
	}

	fromEmail := "noreply@example.com" // Change the sender email address
	fromName := "Your App Name"        // Change the sender name
	subject := "Password Reset Link"

	to := mail.NewEmail("", recipientEmail)
	message := mail.NewSingleEmail(
		mail.NewEmail(fromName, fromEmail),
		subject,
		to,
		"",
		resetLink,
	)

	client := sendgrid.NewSendClient(sendGridAPIKey)

	response, err := client.Send(message)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		return err
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		log.Printf("Email sent successfully, response code: %d\n", response.StatusCode)
		return nil
	}

	log.Printf("Failed to send email, response code: %d\n", response.StatusCode)
	return fmt.Errorf("Failed to send email, response code: %d", response.StatusCode)
}
