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
	gateway.Use(gin.Logger())
	gateway.Use(gin.Recovery())

	// routes definition here
	controller.gateway.GET("/api/searchBiography", controller.searchBiography)
	controller.gateway.GET("/api/searchWork", controller.searchWork)
	controller.gateway.GET("/api/searchKeyword", controller.searchKeyword)
	controller.gateway.GET("/api/searchEmployment", controller.searchEmployment)
	controller.gateway.GET("/api/searchReseacherUrl", controller.searchReseacherUrl)
	controller.gateway.GET("/api/searchEducation", controller.searchEducation)
	controller.gateway.GET("/api/searchEmail", controller.searchEmail)

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
