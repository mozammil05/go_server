// package utils

// import (
// 	"context"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var (
// 	mongoURI           = "mongodb://localhost:27017"
// 	databaseName       = "yourapp"
// 	userCollectionName = "users"
// )

// type Database struct {
// 	Client         *mongo.Client
// 	UserCollection *mongo.Collection
// }

// func InitDB() *Database {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		panic(err)
// 	}

// 	ctx := context.TODO()

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	db := &Database{
// 		Client:         client,
// 		UserCollection: client.Database(databaseName).Collection(userCollectionName),
// 	}

// 	return db
// }

// func (db *Database) Disconnect() {
// 	err := db.Client.Disconnect(context.TODO())
// 	if err != nil {
// 		panic(err)
// 	}
// }

// /
// package utils

// import (
// 	"context"
// 	"log"
// 	"os"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"

// 	"github.com/joho/godotenv"
// )

// type DatabaseConfig struct {
// 	MongoURI           string
// 	DatabaseName       string
// 	UserCollectionName string
// }

// func LoadConfigFromEnv() (*DatabaseConfig, error) {
// 	// Load environment variables from the .env file
// 	if err := godotenv.Load(); err != nil {
// 		return nil, err
// 	}

// 	config := &DatabaseConfig{
// 		MongoURI:           os.Getenv("MONGO_URI"),
// 		DatabaseName:       os.Getenv("DATABASE_NAME"),
// 		UserCollectionName: os.Getenv("USER_COLLECTION_NAME"),
// 	}

// 	return config, nil
// }

// type Database struct {
// 	Client         *mongo.Client
// 	UserCollection *mongo.Collection
// }

// func InitDB() *Database {
// 	// Load database configuration from environment variables
// 	config, err := LoadConfigFromEnv()
// 	if err != nil {
// 		log.Fatal("Error loading database configuration from .env file:", err)
// 	}

// 	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
// 	if err != nil {
// 		panic(err)
// 	}

// 	ctx := context.TODO()

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	db := &Database{
// 		Client:         client,
// 		UserCollection: client.Database(config.DatabaseName).Collection(config.UserCollectionName),
// 	}

// 	return db
// }

// func (db *Database) Disconnect() {
// 	err := db.Client.Disconnect(context.TODO())
// 	if err != nil {
// 		panic(err)
// 	}
// }

// utils/database.go

package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI           = "mongodb://localhost:27017"
	databaseName       = "yourapp"
	userCollectionName = "users"
)

type Database struct {
	Client         *mongo.Client
	UserCollection *mongo.Collection
}

func InitDB() *Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db := &Database{
		Client:         client,
		UserCollection: client.Database(databaseName).Collection(userCollectionName),
	}

	return db
}

func (db *Database) Disconnect() {
	err := db.Client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
