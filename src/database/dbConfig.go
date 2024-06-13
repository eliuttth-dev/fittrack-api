// dbConfig.go

package database

import (
    "context"
    "log"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    client *mongo.Client
    db     *mongo.Database
)

// Connect initializes a connection to MongoDB
func Connect() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://eliuttth-dev-mongodb:Eliuth25012002$@cluster0.35hgh47.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

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
