package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"main/model"
	"net/http"
)

// Gets ORCID biography data from the ORCID API
func (r *Repository) GetWorkData(orcid_id string) (model.OrcidWork, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcid_id+"/works", nil)
	if err != nil {
		return model.OrcidWork{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidWork{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidWork{}, err
	}

	// Transform data into Go Struct
	orcid_work := model.OrcidWork{}
	if err := json.Unmarshal(body, &orcid_work); err != nil {
		return model.OrcidWork{}, err
	}

	return orcid_work, nil
}

func (r *Repository) ProcessOrcidWork(orcid_work model.OrcidWork) model.WorkData {
	work_data := model.WorkData{}
	for _, group := range orcid_work.Group {
		for _, work := range group.WorkSummary {

			var doi string = ""
			if work.ExternalIds.ExternalId != nil && len(work.ExternalIds.ExternalId) > 0 {
				doi = work.ExternalIds.ExternalId[0].Value
			}
			
			count, err := r.GetCitationCount(doi)
			if err != nil {
				fmt.Println("WARNING - could not find citation count for doi " + doi)
			}
		
			work_data.Publications = append(work_data.Publications,
				struct {
					Title   string `json:"title"`
					Doi     string `json:"doi"`
					Url     string `json:"url"`
					Type    string `json:"type"`
					Year    string `json:"year"`
					Journal string `json:"journal"`
					Citations int `json:"citations"`
				}{
					work.Title.Title.Value,
					doi,
					work.Url.Value,
					work.Type,
					work.PublicationDate.Year.Value,
					work.JournalTitle.Value,
					count,
				})
		}
	}

	return work_data
}

