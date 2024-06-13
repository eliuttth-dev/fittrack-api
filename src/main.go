// main.go

package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "fittrack-api/src/routes"
    "fittrack-api/src/database"
)

func main() {
    // Connect to MongoDB
    database.Connect()

    // Retrieve MongoDB client and database
    db := database.GetDatabase()

    // Initialize Gin router
    router := gin.Default()

    // Setup routes with MongoDB dependency
    routes.SetupRoutes(router, db)

    // Run the server
    err := router.Run(":8080")
    if err != nil {
        log.Fatal(err)
    }
}
