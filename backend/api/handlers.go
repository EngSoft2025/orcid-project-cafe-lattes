package api

import (
	"fmt"
	"main/model"
	"net/http"
	"strconv"

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

// too slow....
func (ct * Controller) getHIndexByOrcidId(c * gin.Context) {
	orcidId, ok := c.GetQuery("orcid_id")

	if !ok {
		c.JSON(400, gin.H{"error": "orcid_id is required"})
		return
	}

	err := ct.service.CalculateHIndex(orcidId)
	if err != nil {
		fmt.Println("error: " +err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(200, gin.H{ "message" : "success"})
}

// @Summary Retrieves the citation count given a doi
// @Description Retrieves the citation count given a doi
// @Tags search
// @Accept  json
// @Produce  json
// @Param doi query string true "Work DOI"
// @Success 200 {object} model.CitationCountResponse
// @Router /api/getCitationCount [get]
func (ct * Controller) getCitationCount(c *gin.Context) {
	doi, ok := c.GetQuery("doi")
	if  !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "query must include doi parameter"})
		return
	}

	count, err := ct.service.GetCitationCound(doi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "could not find doi"})
		return
	}

	fmt.Println("doi citation couint: " + strconv.Itoa(count))

	c.JSON(http.StatusOK, model.CitationCountReponse{CitationCount: count})
}
