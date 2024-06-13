// controllers.go

package controllers

import (
    "net/http"
    "time"
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "fittrack-api/src/models"
)

// CreateWorkout handles creation of new workout records
func CreateWorkout(db *mongo.Database) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newWorkout models.Workout

        if err := c.ShouldBindJSON(&newWorkout); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        newWorkout.ID = models.NextID
        models.NextID++
        newWorkout.Date = time.Now()

        collection := db.Collection("workouts")
        result, err := collection.InsertOne(context.TODO(), newWorkout)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Workout added", "insertedID": result.InsertedID})
    }
}
