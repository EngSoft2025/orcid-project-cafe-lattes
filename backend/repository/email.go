package repository

import (
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

func (r *Repository) GetEmailData(orcid_id string) (model.OrcidEmail, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcid_id+"/email", nil)
	if err != nil {
		return model.OrcidEmail{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidEmail{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidEmail{}, err
	}

	// Transform data into Go Struct
	orcid_email := model.OrcidEmail{}
	if err := json.Unmarshal(body, &orcid_email); err != nil {
		return model.OrcidEmail{}, err
	}

	return orcid_email, nil
}

func (r *Repository) ProcessOrcidEmail(orcid_email model.OrcidEmail) model.EmailData {
	email_data := model.EmailData{}

	for _, email := range orcid_email.Emails {
		email_data.Emails = append(email_data.Emails,
			struct {
				Email      string `json:"email"`
				Verified   bool   `json:"verified"`
				Primary    bool   `json:"primary"`
				PutCode    uint64 `json:"put-code"`
				Visibility string `json:"visibility"`
			}{
				email.Email,
				email.Verified,
				email.Primary,
				email.PutCode,
				email.Visibility,
			})
	}

	return email_data
}
