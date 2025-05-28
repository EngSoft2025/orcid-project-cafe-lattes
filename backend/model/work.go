package model

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

type WorkData struct {
	Publications []struct {
		Title   string `json:"title"`
		Doi     string `json:"doi"`
		Url     string `json:"url"`
		Type    string `json:"type"`
		Year    string `json:"year"`
		Journal string `json:"journal"`
	}
}
