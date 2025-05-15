package main

import (
	"fmt"
	"main/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Backend Server")

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	router := gin.Default()
	handlers := handlers.Handlers{}

	router.GET("/api/search", handlers.Search)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":", port)
}
