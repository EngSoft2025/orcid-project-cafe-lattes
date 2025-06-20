package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"main/model"
	"net/http"
	"net/url"
	"strings"
)

func (r *Repository) SearchByName(name string) (model.ResearcherResults, error) {
	splitName := strings.Split(name, " ")
	firstName := splitName[0]
	familyName := strings.Join(splitName[1:], " ")
	fmt.Println("Name:", firstName, familyName)
	// Create new HTTP request
	req, err := createORCIDSearchRequest(r.apiLink, firstName, familyName)
	fmt.Println("Request URL:", req.URL.String())
	if err != nil {
		return model.ResearcherResults{}, fmt.Errorf("error creating request: %w", err)
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.ResearcherResults{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.ResearcherResults{}, err
	}

	researchersIds := model.ResearchersIdsOrcid{}

	if err := json.Unmarshal(body, &researchersIds); err != nil {
		return model.ResearcherResults{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	results, err := r.getOrcidNames(researchersIds)
	if err != nil {
		return model.ResearcherResults{}, fmt.Errorf("error getting ORCID Namess: %w", err)
	}

	return results, nil
}

// createORCIDSearchRequest creates a new, correctly encoded HTTP request for the ORCID API.
func createORCIDSearchRequest(apiLink, firstName, familyName string) (*http.Request, error) {
	// 1. Start with the base URL for the search endpoint.
	// Make sure apiLink ends with a slash, e.g., "https://pub.orcid.org/v3.0/"
	baseURL := apiLink + "search/"

	// 2. Construct the raw query string for the 'q' parameter.
	// We still use Sprintf here just to build the *value*, not the whole URL.
	// Note the quotes around the family name for an exact phrase search.
	orcidQuery := fmt.Sprintf("given-names:%s OR family-name:\"%s\"", firstName, familyName)

	// 3. Create a url.Values map to hold all query parameters.
	params := url.Values{}
	params.Set("q", orcidQuery)
	params.Set("rows", "10")

	// 4. Create the final URL. The Encode() method handles all the necessary
	//    URL encoding for spaces, quotes, and other special characters safely.
	finalURL := baseURL + "?" + params.Encode()

	// 5. Create the HTTP request object.
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *Repository) getOrcidNames(researchersIds model.ResearchersIdsOrcid) (model.ResearcherResults, error) {
	results := model.ResearcherResults{}
	results.NumFound = len(researchersIds.Result)
	for _, researcher := range researchersIds.Result {
		orcidId := researcher.OrcidIdentifier.Path
		researcher, err := r.getOrcidNameById(orcidId)
		if err != nil {
			return model.ResearcherResults{}, fmt.Errorf("error getting ORCID name by ID: %w", err)
		}
		results.Researchers = append(results.Researchers, researcher)
	}
	return results, nil
}

func (r *Repository) getOrcidNameById(orcidId string) (model.Researcher, error) {
	biography, err := r.GetPersonData(orcidId)
	if err != nil {
		return model.Researcher{}, fmt.Errorf("error getting person data: %w", err)
	}

	name := biography.Name.GivenNames.Value + " " + biography.Name.FamilyName.Value

	return model.Researcher{
		Orcid_id: orcidId,
		Name:     name,
	}, nil
}
