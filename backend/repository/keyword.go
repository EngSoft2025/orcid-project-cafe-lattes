package repository

import (
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

func (r *Repository) GetKeywordData(orcid_id string) (model.OrcidKeyword, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcid_id+"/keywords", nil)
	if err != nil {
		return model.OrcidKeyword{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidKeyword{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidKeyword{}, err
	}

	// Transform data into Go Struct
	orcid_keywords := model.OrcidKeyword{}
	if err := json.Unmarshal(body, &orcid_keywords); err != nil {
		return model.OrcidKeyword{}, err
	}

	return orcid_keywords, nil
}

func (r *Repository) ProcessOrcidKeyword(orcid_keyword model.OrcidKeyword) model.KeywordData {
	keyword_data := model.KeywordData{}

	for _, keyword := range orcid_keyword.Keyword {
		keyword_data.Keywords = append(keyword_data.Keywords,
			struct {
				Content    string `json:"content"`
				Visibility string `json:"visibility"`
				PutCode    uint64 `json:"put-code"`
			}{
				keyword.Content,
				keyword.Visibility,
				keyword.PutCode,
			})
	}

	return keyword_data
}
