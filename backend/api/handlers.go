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
		return
	}

	data, err := ct.service.SearchBiography(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " biograpy data"})
		return
	}

	c.JSON(200, data)
}

// Handler that Responds query for Work Data
func (ct *Controller) searchWork(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}

	data, err := ct.service.SearchWork(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " work data"})
		return
	}

	c.JSON(200, data)
}

// @Summary Searches a record by ORCID ID
// @Description Searches a record by ORCID ID
// @Tags record
// @Accept  json
// @Produce  json
// @Param orcid_id query string true "ORCID ID"
// @Success 200 {object} model.RecordDataResponse
// @Router /api/searchRecord [get]
func (ct *Controller) searchRecord(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}

	data, err := ct.service.SearchRecord(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " record data"})
		return
	}

	c.JSON(200, data)
}

// @Summary Searches Researchers by Name
// @Description Searches Researchers by Name
// @Tags search
// @Accept  json
// @Produce  json
// @Param name query string true "Researcher Name"
// @Success 200 {object} model.ResearcherResults
// @Router /api/searchResearchersByName [get]
func (ct *Controller) searchResearchersByName(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.JSON(400, gin.H{"error": "name is required"})
		return
	}
	data, err := ct.service.SearchResearchersByName(name)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get researchers by name: " + name})
		return
	}

	c.JSON(200, data)
}
