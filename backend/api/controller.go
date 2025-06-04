package api

import (
	"fmt"
	"main/service"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "main/docs"
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
	controller.gateway.GET("/api/searchRecord", controller.searchRecord)

	// swagger !!!
	controller.gateway.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
