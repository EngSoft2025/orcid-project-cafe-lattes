package repository

import (
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

func (r *Repository) GetReseacherData(orcid_id string) (model.OrcidResearcherUrl, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcid_id+"/researcher-urls", nil)
	if err != nil {
		return model.OrcidResearcherUrl{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidResearcherUrl{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidResearcherUrl{}, err
	}

	// Transform data into Go Struct
	orcid_reseacher_urls := model.OrcidResearcherUrl{}
	if err := json.Unmarshal(body, &orcid_reseacher_urls); err != nil {
		return model.OrcidResearcherUrl{}, err
	}

	return orcid_reseacher_urls, nil
}

func (r *Repository) ProcessOrcidResearcherUrl(orcid_reseacher_url model.OrcidResearcherUrl) model.ReseacherUrlData {
	reseacher_url_data := model.ReseacherUrlData{}

	for _, reseacher_url := range orcid_reseacher_url.ResearcherUrl {
		reseacher_url_data.Urls = append(reseacher_url_data.Urls,
			struct {
				PutCode    uint64 `json:"put-code"`
				Name       string `json:"name"`
				Value      string `json:"value"`
				Visibility string `json:"visibility"`
			}{
				reseacher_url.PutCode,
				reseacher_url.Url_name,
				reseacher_url.Url.Value,
				reseacher_url.Visibility,
			})
	}

	return reseacher_url_data
}
