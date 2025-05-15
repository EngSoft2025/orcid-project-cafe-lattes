package main

import (
	"fmt"
	"main/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Backend Server")

	router := gin.Default()
	handlers := handlers.Handlers{}

	router.GET("/api/search", handlers.Search)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":", port)
}
