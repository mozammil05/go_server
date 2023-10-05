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
│   ├── main.go
│
├── config/
│   ├── config.go
│   ├── env.go
│
├── controllers/
│   ├── auth_controller.go
│   ├── user_controller.go
│   ├── admin_controller.go
│
├── middleware/
│   ├── auth.go
│   ├── admin_middleware.go
│   ├── superadmin_middleware.go
│
├── models/
│   ├── user.go
│
├── routes/
│   ├── auth_routes.go
│   ├── user_routes.go
│   ├── admin_routes.go
│   ├── routes.go
│
├── services/
│   ├── auth_service.go
│   ├── user_service.go
│   ├── admin_service.go
│
├── utils/
│   ├── jwt.go
│   ├── database.go
│
├── tests/ (optional)
│   ├── integration/
│   ├── unit/
│
├── .env
├── go.mod
├── go.sum
└── README.md

