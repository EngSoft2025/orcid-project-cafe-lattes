package api

import (
	"fmt"
	"main/service"

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
	controller.gateway.GET("/api/search", controller.Search)

	

	return controller
}

func (ct *Controller) Start() error {
	fmt.Println("Running server on port 8080")
	return ct.gateway.Run(":8080") // todo: change to env variable
}
