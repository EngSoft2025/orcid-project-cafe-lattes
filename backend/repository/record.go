package repository

import (
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

func (r *Repository) GetRecordData(orcidID string) (model.OrcidRecord, error) {
	// Create new HTTP request
	req, err := http.NewRequest("GET", r.apiLink+orcidID+"/record", nil)
	if err != nil {
		return model.OrcidRecord{}, err
	}

	// Set the Accept header to request JSON
	req.Header.Set("Accept", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrcidRecord{}, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrcidRecord{}, err
	}

	// Transform data into Go Struct
	orcidRecord := model.OrcidRecord{}
	if err := json.Unmarshal(body, &orcidRecord); err != nil {
		return model.OrcidRecord{}, err
	}

	return orcidRecord, nil
}

func (r *Repository) ProcessOrcidRecord(orcidRecord model.OrcidRecord) model.RecordDataResponse {
	recordData := model.RecordDataResponse{}

	// Get orcid id
	recordData.OrcidID = orcidRecord.OrcidIdentifier.Path

	// Person process
	recordData.Person.GivenName = orcidRecord.Person.Name.GivenNames.Value
	recordData.Person.FamilyName = orcidRecord.Person.Name.FamilyName.Value

	if orcidRecord.Person.Name.CreditName != nil {
		if recordData.Person.CreditName == nil {
			recordData.Person.CreditName = new(string)
		}
		*recordData.Person.CreditName = orcidRecord.Person.Name.CreditName.Value
	}

	if orcidRecord.Person.Biography != nil {
		// Check if BiographyContent itself needs to be initialized
		if recordData.Person.BiographyContent == nil {
			recordData.Person.BiographyContent = new(string) // Allocate memory for the string
		}
		*recordData.Person.BiographyContent = orcidRecord.Person.Biography.Content
	}

	for _, email := range orcidRecord.Person.Emails.Email {
		recordData.Person.Contact.Emails = append(recordData.Person.Contact.Emails, email.Email)
	}
	for _, researcherURL := range orcidRecord.Person.ResearcherURLs.ResearcherURL {
		recordData.Person.Contact.ResearcherURL = append(recordData.Person.Contact.ResearcherURL,
			struct {
				URLName string `json:"url-name"`
				URL     string `json:"url"`
			}{
				researcherURL.URLName,
				researcherURL.URL.Value,
			})
	}

	// Keyword process
	for _, keyword := range orcidRecord.Person.Keywords.Keyword {
		recordData.Keywords = append(recordData.Keywords, keyword.Content)
	}

	// Educations process
	for _, affiliation_group := range orcidRecord.ActivitiesSummary.Educations.AffiliationGroup {
		for _, summaries := range affiliation_group.Summaries {
			var startYear *string
			var startMonth *string
			if summaries.EducationSummary.StartDate != nil {
				startYear = new(string)
				startMonth = new(string)
				*startYear = summaries.EducationSummary.StartDate.Year.Value
				*startMonth = summaries.EducationSummary.StartDate.Month.Value
			}

			var endYear *string
			var endMonth *string
			if summaries.EducationSummary.EndDate != nil {
				endYear = new(string)
				endMonth = new(string)
				*endYear = summaries.EducationSummary.EndDate.Year.Value
				*endMonth = summaries.EducationSummary.EndDate.Month.Value
			}

			var organizationName *string
			if summaries.EducationSummary.Organization != nil {
				organizationName = new(string)
				*organizationName = summaries.EducationSummary.Organization.Name
			}
			var url *string
			if summaries.EducationSummary.URL != nil {
				url = new(string)
				*url = summaries.EducationSummary.URL.Value
			}

			recordData.Affiliations.Eduactions = append(recordData.Affiliations.Eduactions,
				model.AffiliationStruct{
					DepartmentName:   summaries.EducationSummary.DepartmentName,
					RoleTitle:        summaries.EducationSummary.RoleTitle,
					StartYear:        startYear,
					StartMonth:       startMonth,
					EndYear:          endYear,
					EndMonth:         endMonth,
					OrganizationName: organizationName,
					URL:              url,
				})
		}
	}

	// Employment process
	for _, affiliation_group := range orcidRecord.ActivitiesSummary.Employments.AffiliationGroup {
		for _, summaries := range affiliation_group.Summaries {
			var startYear *string
			var startMonth *string
			if summaries.EmploymentSummary.StartDate != nil {
				startYear = new(string)
				startMonth = new(string)
				*startYear = summaries.EmploymentSummary.StartDate.Year.Value
				*startMonth = summaries.EmploymentSummary.StartDate.Month.Value
			}

			var endYear *string
			var endMonth *string
			if summaries.EmploymentSummary.EndDate != nil {
				endYear = new(string)
				endMonth = new(string)
				*endYear = summaries.EmploymentSummary.EndDate.Year.Value
				*endMonth = summaries.EmploymentSummary.EndDate.Month.Value
			}

			var organizationName *string
			if summaries.EmploymentSummary.Organization != nil {
				organizationName = new(string)
				*organizationName = summaries.EmploymentSummary.Organization.Name
			}
			var url *string
			if summaries.EmploymentSummary.URL != nil {
				url = new(string)
				*url = summaries.EmploymentSummary.URL.Value
			}

			recordData.Affiliations.Employments = append(recordData.Affiliations.Employments,
				model.AffiliationStruct{
					DepartmentName:   summaries.EmploymentSummary.DepartmentName,
					RoleTitle:        summaries.EmploymentSummary.RoleTitle,
					StartYear:        startYear,
					StartMonth:       startMonth,
					EndYear:          endYear,
					EndMonth:         endMonth,
					OrganizationName: organizationName,
					URL:              url,
				})
		}
	}

	// Qualifications process
	for _, affiliation_group := range orcidRecord.ActivitiesSummary.Qualifications.AffiliationGroup {
		for _, summaries := range affiliation_group.Summaries {
			var startYear *string
			var startMonth *string
			if summaries.QualificationSummary.StartDate != nil {
				startYear = new(string)
				startMonth = new(string)
				*startYear = summaries.QualificationSummary.StartDate.Year.Value
				*startMonth = summaries.QualificationSummary.StartDate.Month.Value
			}

			var endYear *string
			var endMonth *string
			if summaries.QualificationSummary.EndDate != nil {
				endYear = new(string)
				endMonth = new(string)
				*endYear = summaries.QualificationSummary.EndDate.Year.Value
				*endMonth = summaries.QualificationSummary.EndDate.Month.Value
			}

			var organizationName *string
			if summaries.QualificationSummary.Organization != nil {
				organizationName = new(string)
				*organizationName = summaries.QualificationSummary.Organization.Name
			}
			var url *string
			if summaries.QualificationSummary.URL != nil {
				url = new(string)
				*url = summaries.QualificationSummary.URL.Value
			}

			recordData.Affiliations.Qualifications = append(recordData.Affiliations.Qualifications,
				model.AffiliationStruct{
					DepartmentName:   summaries.QualificationSummary.DepartmentName,
					RoleTitle:        summaries.QualificationSummary.RoleTitle,
					StartYear:        startYear,
					StartMonth:       startMonth,
					EndYear:          endYear,
					EndMonth:         endMonth,
					OrganizationName: organizationName,
					URL:              url,
				})
		}
	}

	// Fundings Process
	for _, group := range orcidRecord.ActivitiesSummary.Fundings.Group {
		for _, summary := range group.FundingSummary {

			var title *string
			if summary.Title != nil {
				title = new(string)
				*title = summary.Title.Title.Value
			}
			var startYear *string
			var startMonth *string
			if summary.StartDate != nil {
				startYear = new(string)
				startMonth = new(string)
				*startYear = summary.StartDate.Year.Value
				*startMonth = summary.StartDate.Month.Value
			}

			var endYear *string
			var endMonth *string
			if summary.StartDate != nil {
				endYear = new(string)
				endMonth = new(string)
				*endYear = summary.EndDate.Year.Value
				*endMonth = summary.EndDate.Month.Value
			}

			var organizationName *string
			if summary.Organization != nil {
				organizationName = new(string)
				*organizationName = summary.Organization.Name
			}

			var url *string
			if summary.URL != nil {
				url = new(string)
				*url = summary.URL.Value
			}

			recordData.Affiliations.Fundings = append(recordData.Affiliations.Fundings,
				model.FundingsStruct{
					Title:            title,
					StartYear:        startYear,
					StartMonth:       startMonth,
					EndYear:          endYear,
					EndMonth:         endMonth,
					OrganizationName: organizationName,
					URL:              url,
				})
		}
	}

	// Works process
	for _, group := range orcidRecord.ActivitiesSummary.Works.Group {
		for _, summary := range group.WorkSummary {

			var title *string
			if summary.Title != nil {
				title = new(string)
				*title = summary.Title.Title.Value
			}
			var doi *string
			if summary.ExternalIDs != nil {
				doi = new(string)
				if len(summary.ExternalIDs.ExternalID) > 1 {
					*doi = summary.ExternalIDs.ExternalID[0].ExternalIDValue
				}
			}
			var url *string
			if summary.URL != nil {
				url = new(string)
				*url = summary.URL.Value
			}
			var year *string
			if summary.PublicationDate != nil {
				year = new(string)
				*year = summary.PublicationDate.Year.Value
			}
			var jounalTitle *string
			if summary.JournalTitle != nil {
				jounalTitle = new(string)
				*jounalTitle = summary.JournalTitle.Value
			}

			recordData.Publications = append(recordData.Publications,
				model.PublicationStruct{
					Title:   title,
					Doi:     doi,
					Url:     url,
					Type:    summary.Type,
					Year:    year,
					Journal: jounalTitle,
				})
		}
	}

	return recordData
}
