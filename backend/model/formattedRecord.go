package model

type AffiliationStruct struct {
	DepartmentName   *string `json:"department-name"`
	RoleTitle        string  `json:"role-title"`
	StartYear        *string `json:"start-year"`
	StartMonth       *string `json:"start-month"`
	EndYear          *string `json:"end-year"`
	EndMonth         *string `json:"end-month"`
	OrganizationName *string `json:"organization-name"`
	URL              *string `json:"url,omitempty"`
}

type PublicationStruct struct {
	Title   *string `json:"title"`
	Doi     *string `json:"doi,omitempty"`
	Url     *string `json:"url,omitempty"`
	Type    string  `json:"type"`
	Year    *string `json:"year"`
	Journal *string `json:"journal,omitempty"`
}

type FundingsStruct struct {
	Title            *string `json:"title"`
	StartYear        *string `json:"start-year"`
	StartMonth       *string `json:"start-month"`
	EndYear          *string `json:"end-year"`
	EndMonth         *string `json:"end-month"`
	OrganizationName *string `json:"organization-name"`
	URL              *string `json:"url,omitempty"`
}

type RecordDataResponse struct {
	OrcidID string `json:"orcid-id"`
	Person  struct {
		GivenName        string  `json:"given-names"`
		FamilyName       string  `json:"family-names"`
		CreditName       *string `json:"credit-names"`
		BiographyContent *string `json:"biography"`
		Contact          struct {
			Emails        []string `json:"email"`
			ResearcherURL []struct {
				URLName string `json:"url-name"`
				URL     string `json:"url"`
			} `json:"researcher-url,omitempty"`
		} `json:"contact"`
	} `json:"person"`

	Keywords []string `json:"keywords"`

	Affiliations struct {
		Eduactions []AffiliationStruct `json:"educations"`

		Qualifications []AffiliationStruct `json:"qualifications"`

		Employments []AffiliationStruct `json:"employments"`

		Fundings []FundingsStruct `json:"fundings"`
	} `json:"affiliations"`

	Publications []PublicationStruct `json:"publications"`

	MetricsData struct {
		PublicationsByYear []struct {
			Year  string `json:"year"`
			Count int    `json:"count"`
		}
		PublicationsByJournal []struct {
			Title string `json:"title"`
			Count int    `json:"count"`
		}
		HIndex int `json:"h-index"`
	}
}
