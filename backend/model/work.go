package model

type (
	WorkData struct{}
)

type OrcidWork struct {
	Group []struct {
		WorkSummary []struct {
			PutCode string `json:"put-code"`
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
