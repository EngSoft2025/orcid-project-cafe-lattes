package model

type RecordDataResponse struct {
	OrcidID string `json:"orcid-id"`
	Person  struct {
		GivenName        string
		FamilyName       string
		CreditName       string
		BiographyContent string
		Contact          struct {
			Emails        []string
			ResearcherURL []struct {
				UrlName string `json:"url-name"`
				Url     struct {
					Value string `json:"value"`
				} `json:"url"`
			} `json:"researcher-url,omitempty"`
		} `json:"contact"`
	} `json:"person"`

	Keywords []string

	Affiliations struct {
		Eduactions []struct {
			DepartmentName   string
			RoleTitle        string
			StartYear        string
			StartMonth       string
			EndYear          string
			EndMonth         string
			OrganizationName string
			URL              string
		}

		Qualifications []struct {
			DepartmentName   string
			RoleTitle        string
			StartYear        string
			StartMonth       string
			EndYear          string
			EndMonth         string
			OrganizationName string
			URL              string
		}

		Employments []struct {
			DepartmentName   string
			RoleTitle        string
			StartYear        string
			StartMonth       string
			EndYear          string
			EndMonth         string
			OrganizationName string
			URL              string
		}

		Fundings []struct {
			Title            string
			StartYear        string
			StartMonth       string
			EndYear          string
			EndMonth         string
			OrganizationName string
			URL              string
		}
	}

	Publications []struct {
		Title   string `json:"title"`
		Doi     string `json:"doi"`
		Url     string `json:"url"`
		Type    string `json:"type"`
		Year    string `json:"year"`
		Journal string `json:"journal"`
	} `json:"publications"`
}
