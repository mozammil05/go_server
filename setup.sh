#!/bin/bash

# Create the main project directory
mkdir -p my-auth-app/handlers my-auth-app/models my-auth-app/routes

# Create the main.go file
cat <<EOL > my-auth-app/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"my-auth-app/routes"
)

func main() {
	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	r := routes.SetupRoutes()
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
EOL

# Create handlers directory and handler files
mkdir -p my-auth-app/handlers
cat <<EOL > my-auth-app/handlers/auth_handler.go
package handlers

import (
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your signup handler logic here
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your login handler logic here
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your change password handler logic here
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your reset password handler logic here
}
EOL

# Create models directory and model files
mkdir -p my-auth-app/models
cat <<EOL > my-auth-app/models/user.go
package models

type User struct {
	ID       int    \`json:"id"\`
	Username string \`json:"username"\`
	Password string \`json:"password"\`
}

// Implement CRUD functions for your user model here
EOL

# Create routes directory and routes file
cat <<EOL > my-auth-app/routes/routes.go
package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"my-auth-app/handlers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/changepassword", handlers.ChangePasswordHandler).Methods("POST")
	r.HandleFunc("/resetpassword", handlers.ResetPasswordHandler).Methods("POST")

	return r
}
EOL

# Create go.mod file
cat <<EOL > my-auth-app/go.mod
module my-auth-app

go 1.17

require (
	github.com/gorilla/mux v1.8.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/dgrijalva/jwt-go v3.2.0
)
EOL

echo "Authentication project structure and sample files created successfully!"
