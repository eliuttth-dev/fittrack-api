package routes

import (
	"github.com/gin-gonic/gin"
	"fittrack-api/src/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	router.POST("/v1/api/workouts", controllers.CreateWorkout)

	return router
}
