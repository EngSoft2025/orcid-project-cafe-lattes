package api

import (
	"github.com/gin-gonic/gin"
)

// handlers here (they are able to call the service layer through ct.service.call_something())

// Handler that Responds query for Biographical Data
func (ct *Controller) searchBiography(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
	}

	data, err := ct.service.SearchBiography(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " biograpy data"})
	}

	c.JSON(200, data)
}
