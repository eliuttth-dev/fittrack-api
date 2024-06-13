// models.go

package models

import "time"

// Workout represents a workout model
type Workout struct {
    ID          int       `json:"id" bson:"id"`
    Title       string    `json:"title" bson:"title"`
    Description string    `json:"description" bson:"description"`
    Date        time.Time `json:"date" bson:"date"`
}

var Workouts []Workout
var NextID = 1
