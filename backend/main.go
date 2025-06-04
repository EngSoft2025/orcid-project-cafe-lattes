package main

import (
	"fmt"
	"main/api"
	"main/repository"
	"main/service"
	"net/http"

	"github.com/joho/godotenv"
)

// @title ORCID API
// @version 1.0.0
// @contact.name Caffe Lattes
// @host localhost:port
// @BasePath /

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permitir todas as origens
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

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
