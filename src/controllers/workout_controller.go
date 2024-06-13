package controllers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"fittrack-api/src/models"
)

func CreateWorkout(c *gin.Context){
	var newWorkout models.Workout

	if err := c.ShouldBindJSON(&newWorkout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newWorkout.ID = models.NextID
	models.NextID++
	newWorkout.Date = time.Now()

	models.Workouts = append(models.Workouts, newWorkout)

	c.JSON(http.StatusCreated, newWorkout)
}
