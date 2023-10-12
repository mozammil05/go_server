package models

import "time"

type User struct {
	Email        string    `json:"email" binding:"omitempty"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	IsActive     bool      `json:"is_active"`
	Expiration   time.Time `bson:"expiration"`
	Tokens       string    `bson:"tokens"`
	Created      time.Time `bson:"created"`
	Updated      time.Time `bson:"updated"`
	ProfileImage string    `json:"profileImage"`
}

type ChangePasswordInput struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password" `
	NewPassword string `json:"new_password"`
}

// package models

// type User struct {
// 	Email    string `json:"email" binding:"required,email,regex=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,4}$" meta:"email is required and must be a valid email address"`
// 	Username string `json:"username" binding:"required,min=5,max=50" meta:"username is required and must be between 5 and 50 characters"`
// 	Password string `json:"password" binding:"required,min=6" meta:"password is required and must be at least 6 characters"`
// }

// models/user.go

// package models

// type User struct {
// 	Email    string `json:"email" binding:"required,customEmail" meta:"email is required and must be a valid email address"`
// 	Username string `json:"username" binding:"required,min=5,max=50" meta:"username is required and must be between 5 and 50 characters"`
// 	Password string `json:"password" binding:"required,customPassword,min=6" meta:"password is required and must be at least 6 characters"`
// }
