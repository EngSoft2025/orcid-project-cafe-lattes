package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func recoverMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(500, gin.H{
				"error": r,
			})
		}
	}()
	c.Next()
}

// handlers here (they are able to call the service layer through ct.service.call_something())

// Handler that Responds query for Biographical Data
func (ct *Controller) searchBiography(c *gin.Context) {
	defer recoverMiddleware(c)

	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		log.Panicln("Query does not contain orcid_id")
	}

	data := ct.service.SearchBiography(orcid_id)

	c.JSON(200, data)
}
