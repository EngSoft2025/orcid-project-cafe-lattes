package model

type ResearchersIdsOrcid struct {
	Result []struct {
		OrcidIdentifier struct {
			Path string `json:"path"`
		} `json:"orcid-identifier"`
	} `json:"result"`
	NumFound int `json:"num-found"`
}

type Researcher struct {
	Orcid_id string `json:"orcid_id"`
	Name     string `json:"name"`
}

type ResearcherResults struct {
	Researchers []Researcher `json:"researchers"`
	NumFound    int          `json:"num_found"`
}
