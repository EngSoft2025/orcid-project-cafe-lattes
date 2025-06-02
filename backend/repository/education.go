package repository

import (
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

func (r *Repository) GetEducationData(orcid_id string) (model.OrcidEducation, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcid_id+"/educations", nil)
	if err != nil {
		return model.OrcidEducation{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidEducation{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidEducation{}, err
	}

	// Transform data into Go Struct
	orcid_education := model.OrcidEducation{}
	if err := json.Unmarshal(body, &orcid_education); err != nil {
		return model.OrcidEducation{}, err
	}

	return orcid_education, nil
}

func (r *Repository) ProcessOrcidEducation(orcid_education model.OrcidEducation) model.EducationData {
	education_data := model.EducationData{}

	for _, affiliation := range orcid_education.Affiliation_group {
		for _, summary := range affiliation.Summaries {
			education_data.Education = append(education_data.Education,
				struct {
					Department_name      string `json:"department-name"`
					Role_title           string `json:"role-title"`
					Organization_name    string `json:"organization-name"`
					Organization_address struct {
						City    string `json:"city"`
						Region  string `json:"region"`
						Country string `json:"country"`
					} `json:"organization-address"`
					Start_year   string `json:"start-year"`
					End_year     string `json:"end-year"`
					Url          string `json:"url"`
					External_ids string `json:"external-ids"`
				}{
					summary.EducationSummary.Department_name,
					summary.EducationSummary.Role_title,
					summary.EducationSummary.Organization.Name,
					summary.EducationSummary.Organization.Addess,
					summary.EducationSummary.Start_date.Year.Value,
					summary.EducationSummary.End_date.Year.Value,
					summary.EducationSummary.Url,
					summary.EducationSummary.ExternalIDs,
				})
		}
	}

	return education_data
}
