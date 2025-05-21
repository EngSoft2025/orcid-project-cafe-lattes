package model

// transaction model data here for dto

// DataModel defined in OrcidPublicAPI
type OrcidBiography struct {
	Name struct {
		GivenNames struct {
			Value string `json:"value"`
		} `json:"given-names"`
		FamilyName struct {
			Value string `json:"value"`
		} `json:"family-name"`
		Path string `json:"path"`
	} `json:"name"`
	OtherNames struct {
		OtherName []any `json:"other-name"`
	} `json:"other-names"`
	Biography struct {
		Content string `json:"content"`
	} `json:"biography"`
	ResearcherUrls struct {
		ResearcherURL []struct {
			URL struct {
				Value string `json:"value"`
			} `json:"url"`
		} `json:"researcher-url"`
	} `json:"researcher-urls"`
	Emails struct {
		Email []string `json:"email"`
	} `json:"emails"`
	Addresses struct {
		Address []struct {
			Country struct {
				Value string `json:"value"`
			} `json:"country"`
		} `json:"address"`
	} `json:"addresses"`
	Keywords struct {
		Keyword []struct {
			Content string `json:"content"`
		} `json:"keyword"`
	} `json:"keywords"`
	ExternalIdentifiers struct {
		ExternalIdentifier []struct {
			ExternalIDType  string `json:"external-id-type"`
			ExternalIDValue string `json:"external-id-value"`
			ExternalIDURL   struct {
				Value string `json:"value"`
			} `json:"external-id-url"`
			ExternalIDRelationship string `json:"external-id-relationship"`
		} `json:"external-identifier"`
	} `json:"external-identifiers"`
}

// BiographyDataModel to be sent to frontend client
type BiographyData struct {
	Name      string `json:"name"`
	OrcidId   string `json:"orcid_id"`
	Site      string `json:"site"`
	Biography string `json:"biography"`

	Emails         []string `json:"emails"`
	ResearcherUrls []string `json:"researcher_urls"`
	Country        []string `json:"countries"`
	Keywords       []string `json:"keywords"`

	ExternalIDs []struct {
		ExternalIDType  string `json:"external-id-type"`
		ExternalIDValue string `json:"external-id-value"`
		ExternalIDURL   struct {
			Value string `json:"value"`
		} `json:"external-id-url"`
		ExternalIDRelationship string `json:"external-id-relationship"`
	} `json:"external-identifier"`
}
