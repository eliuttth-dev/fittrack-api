package models

import "time"

type Workout struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Date time.Time `json:"date"`
}

var Workouts = []Workout{}
var NextID = 1