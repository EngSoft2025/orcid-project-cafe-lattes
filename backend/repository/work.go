package repository

import (
	"encoding/json"
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

	return work_data
}
