#!/bin/bash

# Create the main project directory
mkdir -p my-auth-app/cmd my-auth-app/config my-auth-app/controllers my-auth-app/middleware my-auth-app/models my-auth-app/routes my-auth-app/services my-auth-app/utils my-auth-app/tests/integration my-auth-app/tests/unit my-auth-app/docs

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
	r := routes.NewRouter()
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
EOL

# Create config directory and config files
mkdir -p my-auth-app/config
touch my-auth-app/config/config.go my-auth-app/config/env.go

# Create controllers directory and controller files
mkdir -p my-auth-app/controllers
touch my-auth-app/controllers/auth_controller.go my-auth-app/controllers/user_controller.go my-auth-app/controllers/admin_controller.go

# Create middleware directory and middleware files
mkdir -p my-auth-app/middleware
touch my-auth-app/middleware/auth.go my-auth-app/middleware/admin_middleware.go my-auth-app/middleware/superadmin_middleware.go

# Create models directory and model files
mkdir -p my-auth-app/models
touch my-auth-app/models/user.go

# Create routes directory and route files
mkdir -p my-auth-app/routes
touch my-auth-app/routes/auth_routes.go my-auth-app/routes/user_routes.go my-auth-app/routes/admin_routes.go my-auth-app/routes/routes.go

# Create services directory and service files
mkdir -p my-auth-app/services
touch my-auth-app/services/auth_service.go my-auth-app/services/user_service.go my-auth-app/services/admin_service.go

# Create utils directory and utility files
mkdir -p my-auth-app/utils
touch my-auth-app/utils/jwt.go my-auth-app/utils/database.go

# Create tests directory and test files
mkdir -p my-auth-app/tests/integration
mkdir -p my-auth-app/tests/unit

# Create .env file
touch my-auth-app/.env

# Create go.mod file
cat <<EOL > my-auth-app/go.mod
module my-auth-app

go 1.17

require (
    github.com/gin-gonic/gin v1.7.4
    github.com/gorilla/mux v1.8.0
    github.com/golang-jwt/jwt v1.0.2
)
EOL

# Create README.md file
touch my-auth-app/README.md

# Create docs directory and Swagger-related files
mkdir -p my-auth-app/docs
touch my-auth-app/docs/docs.go my-auth-app/docs/swagger.json

# Create swag-init.go file (if needed for Swagger)
touch my-auth-app/swag-init.go

echo "Authentication project structure and sample files created successfully!"
