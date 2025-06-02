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

func (ct *Controller) searchKeyword(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}
	data, err := ct.service.SearchKeyword(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " keyword data"})
		return
	}

	c.JSON(200, data)
}

func (ct *Controller) searchEmployment(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}
	data, err := ct.service.SearchEmployment(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " employment data"})
		return
	}

	c.JSON(200, data)
}

func (ct *Controller) searchReseacherUrl(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}
	data, err := ct.service.SearchReseacherUrls(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " reseacher url data"})
		return
	}

	c.JSON(200, data)
}

func (ct *Controller) searchEducation(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}
	data, err := ct.service.SearchEducation(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " education data"})
		return
	}

	c.JSON(200, data)
}

func (ct *Controller) searchEmail(c *gin.Context) {
	orcid_id, ok := c.GetQuery("orcid_id")
	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}
	data, err := ct.service.SearchEmail(orcid_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not get " + orcid_id + " email data"})
		return
	}

	c.JSON(200, data)
}
