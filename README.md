# go_server
<!-- for create folder this script -->

<!-- chmod +x setup.sh
./setup.sh -->
<!-- server reload use this:=> nodemon --exec go run main.go --signal SIGTERM -->

<!-- for email and password for custom validator -->
<!-- go get -u github.com/go-playground/validator/v10 -->


├── main.go
├── routes
│   └── routes.go
├── controllers
│   ├── auth.go
│   └── user.go
├── models
│   └── user.go
├── utils
│   ├── database.go
│   └── jwt.go
└── middleware
    └── auth.go
