package repository

import (
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

// all the data fetching and treatment logic here
// its called from the service layer
//

type Repository struct {
	apiLink string
	// relevant fields here, maybe orcid client, etc
}

func NewRepository() *Repository {
	return &Repository{apiLink: "https://pub.orcid.org/v3.0/"}
}

// Gets ORCID biography data from the ORCID API
func (r *Repository) GetPersonData(orcid_id string) (model.OrcidBiography, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcid_id+"/person", nil)
	if err != nil {
		return model.OrcidBiography{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidBiography{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidBiography{}, err
	}

	// Transform data into Go Struct
	orcid_biography := model.OrcidBiography{}
	if err := json.Unmarshal(body, &orcid_biography); err != nil {
		return model.OrcidBiography{}, err
	}

	return orcid_biography, nil
}

// Processes the ORCID biography data into a more usable format for use in frontend
func (r *Repository) ProcessOrcidBiography(orcid_biography model.OrcidBiography) model.BiographyData {
	biography_data := model.BiographyData{}

	biography_data.Name = orcid_biography.Name.GivenNames.Value + " " + orcid_biography.Name.FamilyName.Value
	biography_data.OrcidId = orcid_biography.Name.Path
	biography_data.Biography = orcid_biography.Biography.Content
	biography_data.Emails = orcid_biography.Emails.Email

	for _, research_url := range orcid_biography.ResearcherUrls.ResearcherURL {
		biography_data.ResearcherUrls = append(biography_data.ResearcherUrls, research_url.URL.Value)
	}
	for _, address := range orcid_biography.Addresses.Address {
		biography_data.Country = append(biography_data.Country, address.Country.Value)
	}
	for _, keystruct := range orcid_biography.Keywords.Keyword {
		biography_data.Keywords = append(biography_data.Keywords, keystruct.Content)
	}
	biography_data.ExternalIDs = orcid_biography.ExternalIdentifiers.ExternalIdentifier

	return biography_data
}
