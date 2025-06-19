package model

type OrcidWork struct {
	Group []struct {
		WorkSummary []struct {
			PutCode uint64 `json:"put-code"`
			Title   struct {
				Title struct {
					Value string `json:"value"`
				} `json:"title"`
			} `json:"title"`
			ExternalIds struct {
				ExternalId []struct {
					Value string `json:"external-id-value"`
				} `json:"external-id"`
			} `json:"external-ids"`
			Url struct {
				Value string `json:"value"`
			} `json:"url"`
			Type            string `json:"type"`
			PublicationDate struct {
				Year struct {
					Value string `json:"value"`
				} `json:"year"`
			} `json:"publication-date"`
			JournalTitle struct {
				Value string `json:"value"`
			} `json:"journal-title"`
		} `json:"work-summary"`
	} `json:"group"`
}

type WorkData struct {
	Publications []struct {
		Title     string `json:"title"`
		Doi       string `json:"doi"`
		Url       string `json:"url"`
		Type      string `json:"type"`
		Year      string `json:"year"`
		Journal   string `json:"journal"`
		Citations int    `json:"citations"`
	} `json:"publications"`
}

// h index
type HExternalID struct {
	Type  string `json:"external-id-type"`
	Value string `json:"external-id-value"`
}

type HWorkSummary struct {
	ExternalIDs struct {
		ExternalID []HExternalID `json:"external-id"`
	} `json:"external-ids"`
}

type HWorksResponse struct {
	Group []struct {
		WorkSummary []HWorkSummary `json:"work-summary"`
	} `json:"group"`
}

// citation couint

type Paper struct {
	CitationCount int `json:"citationCount"`
}

type CitationCountReponse struct {
	CitationCount int `json:citation-count"`
}
