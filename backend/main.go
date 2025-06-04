package main

import (
	"fmt"
	"main/api"
	"main/repository"
	"main/service"

	"github.com/joho/godotenv"
)

// @title ORCID API
// @version 1.0.0
// @contact.name Caffe Lattes
// @host localhost:port
// @BasePath /
func main() {
	fmt.Println("Starting Backend Server")

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	repository := repository.NewRepository()

	service := service.NewService(repository)

	controller := api.NewController(service)

	fmt.Println("Controller successfully created. Starting server...")
	
	if err := controller.Start(); err != nil {
		fmt.Println("Error starting controller:", err)
	}
}
