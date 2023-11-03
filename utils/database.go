package utils

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define package-level variables here
var (
	// MongoDB connection string
	mongoURI = getEnv("MONGO_URI", "mongodb://localhost:27017")

	// Database name
	databaseName = getEnv("DATABASE_NAME", "yourapp")

	// User collection
	// userCollection    = getEnv("USER_COLLECTION", "users")
	// productCollection = getEnv("PRODUCT_COLLECTION", "products")
	userCollection    = "users"
	productCollection = "products"
)

// getEnv is a helper function to read an environment variable or provide a default value.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// Database represents the MongoDB database and collections.
type Database struct {
	Client            *mongo.Client
	UserCollection    *mongo.Collection
	ProductCollection *mongo.Collection // Add more collection fields as needed
	// Add additional collection fields here
}

// InitDB initializes the database connection.
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
		Client:            client,
		UserCollection:    client.Database(databaseName).Collection(userCollection),
		ProductCollection: client.Database(databaseName).Collection(productCollection),
		// Assign more collections as needed
	}

	return db
}

// Disconnect closes the database connection.
func (db *Database) Disconnect() {
	err := db.Client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
