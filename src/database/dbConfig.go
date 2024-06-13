// dbConfig.go

package database

import (
    "context"
    "log"
    "fmt"
	"os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
)

var (
    client *mongo.Client
    db     *mongo.Database
)

// Connect initializes a connection to MongoDB
func Connect() {

	errorEnv := godotenv.Load()

	if errorEnv != nil {
		log.Fatalf("Error loading .env file: %v", errorEnv)
	}

	dbHost := os.Getenv("DB_HOST")

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s", dbHost))

    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    db = client.Database("Cluster0")
    fmt.Println("Connected to MongoDB!")
}

// GetClient returns the MongoDB client instance
func GetClient() *mongo.Client {
    return client
}

// GetDatabase returns the MongoDB database instance
func GetDatabase() *mongo.Database {
    return db
}
