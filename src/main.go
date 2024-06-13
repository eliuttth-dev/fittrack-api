package main

import (
	"fittrack-api/src/routes"
)

func main(){
	router := routes.SetupRouter()
	router.Run(":8080")
}