// routes.go

package routes

import (
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "fittrack-api/src/controllers"
)

// SetupRoutes initializes routes and handlers
func SetupRoutes(router *gin.Engine, db *mongo.Database) {
    v1 := router.Group("/v1/api")

    // Example route using MongoDB
    v1.POST("/workouts", controllers.CreateWorkout(db))
}
