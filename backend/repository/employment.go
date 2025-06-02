package repository

import (
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

func (r *Repository) GetEmploymentData(orcid_id string) (model.OrcidEmployment, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcid_id+"/employments", nil)
	if err != nil {
		return model.OrcidEmployment{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidEmployment{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidEmployment{}, err
	}

	// Transform data into Go Struct
	orcid_employments := model.OrcidEmployment{}
	if err := json.Unmarshal(body, &orcid_employments); err != nil {
		return model.OrcidEmployment{}, err
	}

	return orcid_employments, nil
}

func (r *Repository) ProcessOrcidEmployment(orcid_employment model.OrcidEmployment) model.EmploymentData {
	employment_data := model.EmploymentData{}

	for _, affiliation := range orcid_employment.Affiliation_group {
		for _, summary := range affiliation.Summaries {
			employment_data.Employment = append(employment_data.Employment,
				struct {
					Department_name      string `json:"department-name"`
					Role_title           string `json:"role-title"`
					Organization_name    string `json:"organization-name"`
					Organization_address struct {
						City    string `json:"city"`
						Region  string `json:"region"`
						Country string `json:"country"`
					} `json:"organization-address"`
					Start_year string `json:"start-year"`
					End_year   string `json:"end-year"`
				}{
					summary.EmploymentSummary.Department_name,
					summary.EmploymentSummary.Role_title,
					summary.EmploymentSummary.Organization.Name,
					summary.EmploymentSummary.Organization.Addess,
					summary.EmploymentSummary.Start_date.Year.Value,
					summary.EmploymentSummary.End_date.Year.Value,
				})
		}
	}

	return employment_data
}
