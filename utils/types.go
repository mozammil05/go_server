package utils

import "time"

// SignResponse represents the response structure for sign operations.
type SignResponse struct {
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Created  time.Time `bson:"created"`
	Updated  time.Time `bson:"updated"`
}

// User represents user information.
type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// UserResponse represents the user response structure.
type UserResponse struct {
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Expiration time.Time `bson:"expiration"`
}

// Token represents a user's token information.
type Token struct {
	Email      string    `json:"email"`
	IsActive   bool      `json:"is_active"`
	Expiration time.Time `json:"expiration"`
	Tokens     string    `json:"tokens"`
}
type UserProfileResponse struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	ProfileImage string `json:"profileImage"`
}

type UpdateResponse struct {
	Username     string `json:"username"`
	ProfileImage string `json:"profileImage"`
}
