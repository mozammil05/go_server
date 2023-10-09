# go_server
<!-- for create folder this script -->

<!-- chmod +x setup.sh
./setup.sh -->
<!-- server reload use this:=> nodemon --exec go run main.go --signal SIGTERM -->

<!-- for email and password for custom validator -->
<!-- go get -u github.com/go-playground/validator/v10 -->


my-auth-app/
│
├── cmd/
│   ├── main.go             // Application entry point
│
├── config/
│   ├── config.go           // Configuration handling
│   ├── env.go              // Environment variable setup
│
├── controllers/
│   ├── auth_controller.go   // Authentication related handlers
│   ├── user_controller.go   // User management handlers
│   ├── admin_controller.go  // Admin-related handlers
│
├── middleware/
│   ├── auth.go             // Authentication middleware
│   ├── admin_middleware.go // Admin middleware
│   ├── superadmin_middleware.go // Superadmin middleware
│
├── models/
│   ├── user.go             // User data model
│
├── routes/
│   ├── auth_routes.go      // Authentication routes
│   ├── user_routes.go      // User management routes
│   ├── admin_routes.go     // Admin routes
│   ├── routes.go           // Router setup
│
├── services/
│   ├── auth_service.go     // Authentication service
│   ├── user_service.go     // User management service
│   ├── admin_service.go    // Admin service
│   ├── email_service.go    // Email service (for sending reset emails)
│
├── utils/
│   ├── jwt.go              // JWT (JSON Web Token) handling
│   ├── database.go         // Database setup
│
├── tests/ (optional)
│   ├── integration/        // Integration tests
│   ├── unit/               // Unit tests
│
├── .env                    // Environment configuration file
├── go.mod                  // Go module file
├── go.sum                  // Go module checksum file
└── README.md               // Project documentation


