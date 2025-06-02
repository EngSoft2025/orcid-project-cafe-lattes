package model

type OrcidEducation struct {
	Affiliation_group []struct {
		Summaries []struct {
			EducationSummary struct {
				PutCode         uint64 `json:"put-code"`
				Department_name string `json:"department-name"`
				Role_title      string `json:"role-title"`

				Start_date struct {
					Year struct {
						Value string `json:"value"`
					} `json:"year"`

					Month struct {
						Value string `json:"value"`
					} `json:"month"`

					Day struct {
						Value string `json:"value"`
					} `json:"day"`
				} `json:"start-date"`

				End_date struct {
					Year struct {
						Value string `json:"value"`
					} `json:"year"`

					Month struct {
						Value string `json:"value"`
					} `json:"month"`

					Day struct {
						Value string `json:"value"`
					} `json:"day"`
				} `json:"end-date"`

				Organization struct {
					Name   string `json:"name"`
					Addess struct {
						City    string `json:"city"`
						Region  string `json:"region"`
						Country string `json:"country"`
					} `json:"address"`
				} `json:"organization"`

				Url string `json:"url"`

				ExternalIDs string `json:"external-ids"`
			} `json:"education-summary"`
		} `json:"summaries"`
	} `json:"affiliation-group"`
}

type EducationData struct {
	Education []struct {
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
	} `json:"education"`
}
