package model

type OrcidResearcherUrl struct {
	ResearcherUrl []struct {
		PutCode  uint64 `json:"put-code"`
		Url_name string `json:"url-name"`
		Url      struct {
			Value string `json:"value"`
		} `json:"url"`
		Visibility string `json:"visibility"`
	} `json:"researcher-url"`
}

type ReseacherUrlData struct {
	Urls []struct {
		PutCode    uint64 `json:"put-code"`
		Name       string `json:"name"`
		Value      string `json:"value"`
		Visibility string `json:"visibility"`
	} `json:"url"`
}
