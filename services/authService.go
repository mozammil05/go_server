// services/auth_service.go

package services

import (
	"context"
	"my-auth-app/utils"

	"go.mongodb.org/mongo-driver/bson"
)

// GetUserEmailByResetToken retrieves the user's email based on the reset token
func GetUserEmailByResetToken(resetToken string) (string, error) {
	// Establish a database connection (assuming you have a db variable in utils.Database)
	db := utils.InitDB()

	// Define the filter to find the user with the provided reset token
	filter := bson.M{"reset_token": resetToken}

	// Define a struct to store the user data
	var user struct {
		Email string `bson:"email"`
	}

	// Query the database to find the user by reset token
	err := db.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
