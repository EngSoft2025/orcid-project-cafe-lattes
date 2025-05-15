package api

import "github.com/gin-gonic/gin"

// handlers here (they are able to call the service layer through ct.service.call_something())

func (ct *Controller) Search(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
