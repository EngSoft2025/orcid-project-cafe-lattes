package api

import (
	"fmt"
	"main/service"
	"net/http"
	"os"

	_ "main/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// CORS middleware
	gateway.Use(func(c *gin.Context) {
		// Configura os cabeçalhos de CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Responde às requisições preflight (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Certifique-se de incluir este cabeçalho
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// other middleware definitions
	gateway.Use(gin.Logger())
	gateway.Use(gin.Recovery())

	// routes definition
	controller.gateway.GET("/api/searchBiography", controller.searchBiography)
	controller.gateway.GET("/api/searchWork", controller.searchWork)
	controller.gateway.GET("/api/searchRecord", controller.searchRecord)
	controller.gateway.GET("/api/searchResearchersByName", controller.searchResearchersByName)
	controller.gateway.GET("/api/getCitationCount", controller.getCitationCount)

	// too slow for usage, test purpose only
	controller.gateway.GET("/api/getHIndex", controller.getHIndexByOrcidId)

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
	return ct.gateway.Run(":" + port)
}
