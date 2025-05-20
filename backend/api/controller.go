package api

import (
	"fmt"
	"main/service"
	"os"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	gateway *gin.Engine
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	gateway := gin.New()

	controller := &Controller{
		gateway: gateway,
		service: service,
	}

	// middleware definitions here

	// routes definition here
	controller.gateway.GET("/api/searchBiography", controller.searchBiography)

	return controller
}

func (ct *Controller) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	fmt.Println("Running server on port", port)
	return ct.gateway.Run(":" + port) // todo: change to env variable
}
