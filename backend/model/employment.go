package model

type OrcidEmployment struct {
	Affiliation_group []struct {
		Summaries []struct {
			EmploymentSummary struct {
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
			} `json:"employment-summary"`
		} `json:"summaries"`
	} `json:"affiliation-group"`
}

type EmploymentData struct {
	Employment []struct {
		Department_name      string `json:"department-name"`
		Role_title           string `json:"role-title"`
		Organization_name    string `json:"organization-name"`
		Organization_address struct {
			City    string `json:"city"`
			Region  string `json:"region"`
			Country string `json:"country"`
		} `json:"organization-address"`
		Start_year string `json:"start-year"`
		End_year   string `json:"end-year"`
	} `json:"employments"`
}
