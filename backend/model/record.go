package model

// Estruturas auxiliares comuns
type StringValue struct {
	Value string `json:"value"`
}

type Int64Value struct {
	Value int64 `json:"value"`
}

type OrcidIdentifierInfo struct {
	URI  string `json:"uri,omitempty"`
	Path string `json:"path,omitempty"`
	Host string `json:"host,omitempty"`
}

// Estrutura principal
type OrcidRecord struct {
	OrcidIdentifier   *OrcidIdentifierInfo `json:"orcid-identifier,omitempty"`
	Preferences       *Preferences         `json:"preferences,omitempty"`
	History           *History             `json:"history,omitempty"`
	Person            *Person              `json:"person,omitempty"`
	ActivitiesSummary *ActivitiesSummary   `json:"activities-summary,omitempty"`
	Path              string               `json:"path,omitempty"`
}

type Preferences struct {
	Locale string `json:"locale,omitempty"`
}

type History struct {
	CreationMethod       string      `json:"creation-method,omitempty"`
	CompletionDate       *Int64Value `json:"completion-date,omitempty"`
	SubmissionDate       *Int64Value `json:"submission-date,omitempty"`
	LastModifiedDate     *Int64Value `json:"last-modified-date,omitempty"`
	Claimed              bool        `json:"claimed"`
	Source               *Source     `json:"source,omitempty"`
	DeactivationDate     *Int64Value `json:"deactivation-date,omitempty"`
	VerifiedEmail        bool        `json:"verified-email"`
	VerifiedPrimaryEmail bool        `json:"verified-primary-email"`
}

type Person struct {
	LastModifiedDate    *Int64Value                `json:"last-modified-date,omitempty"`
	Name                *Name                      `json:"name,omitempty"`
	OtherNames          *OtherNames                `json:"other-names,omitempty"`
	Biography           *Biography                 `json:"biography,omitempty"`
	ResearcherURLs      *ResearcherURLs            `json:"researcher-urls,omitempty"`
	Emails              *Emails                    `json:"emails,omitempty"`
	Addresses           *Addresses                 `json:"addresses,omitempty"`
	Keywords            *Keywords                  `json:"keywords,omitempty"`
	ExternalIdentifiers *PersonExternalIdentifiers `json:"external-identifiers,omitempty"`
	Path                string                     `json:"path,omitempty"`
}

type Name struct {
	CreatedDate      *Int64Value  `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value  `json:"last-modified-date,omitempty"`
	GivenNames       *StringValue `json:"given-names,omitempty"`
	FamilyName       *StringValue `json:"family-name,omitempty"`
	CreditName       *StringValue `json:"credit-name,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Visibility       string       `json:"visibility,omitempty"`
	Path             string       `json:"path,omitempty"`
}

type OtherNames struct {
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	OtherName        []OtherName `json:"other-name,omitempty"` // Estrutura de OtherName não detalhada, pode precisar de ajuste
	Path             string      `json:"path,omitempty"`
}

// OtherName: Se tiver uma estrutura específica, defina aqui. Por enquanto, está vazio no JSON.
type OtherName struct {
	// Campos de OtherName se existirem
}

type Biography struct {
	CreatedDate      *Int64Value `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	Content          string      `json:"content,omitempty"`
	Visibility       string      `json:"visibility,omitempty"`
	Path             string      `json:"path,omitempty"`
}

type ResearcherURLs struct {
	LastModifiedDate *Int64Value     `json:"last-modified-date,omitempty"`
	ResearcherURL    []ResearcherURL `json:"researcher-url,omitempty"`
	Path             string          `json:"path,omitempty"`
}

type ResearcherURL struct {
	CreatedDate      *Int64Value  `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value  `json:"last-modified-date,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	URLName          string       `json:"url-name,omitempty"`
	URL              *StringValue `json:"url,omitempty"`
	Visibility       string       `json:"visibility,omitempty"`
	Path             string       `json:"path,omitempty"`
	PutCode          int64        `json:"put-code"`
	DisplayIndex     int          `json:"display-index"`
}

type Emails struct {
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	Email            []Email     `json:"email,omitempty"`
	Path             string      `json:"path,omitempty"`
}

type Email struct {
	CreatedDate      *Int64Value `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	Source           *Source     `json:"source,omitempty"`
	Email            string      `json:"email,omitempty"`
	Path             *string     `json:"path,omitempty"`
	Visibility       string      `json:"visibility,omitempty"`
	Verified         bool        `json:"verified"`
	Primary          bool        `json:"primary"`
	PutCode          *int64      `json:"put-code,omitempty"`
}

type Addresses struct {
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	Address          []Address   `json:"address,omitempty"`
	Path             string      `json:"path,omitempty"`
}

type Address struct {
	CreatedDate      *Int64Value  `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value  `json:"last-modified-date,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Country          *StringValue `json:"country,omitempty"`
	Visibility       string       `json:"visibility,omitempty"`
	Path             string       `json:"path,omitempty"`
	PutCode          int64        `json:"put-code"`
	DisplayIndex     int          `json:"display-index"`
}

type Keywords struct {
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	Keyword          []Keyword   `json:"keyword,omitempty"`
	Path             string      `json:"path,omitempty"`
}

type Keyword struct {
	CreatedDate      *Int64Value `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	Source           *Source     `json:"source,omitempty"`
	Content          string      `json:"content,omitempty"`
	Visibility       string      `json:"visibility,omitempty"`
	Path             string      `json:"path,omitempty"`
	PutCode          int64       `json:"put-code"`
	DisplayIndex     int         `json:"display-index"`
}

type PersonExternalIdentifiers struct {
	LastModifiedDate   *Int64Value                `json:"last-modified-date,omitempty"`
	ExternalIdentifier []ExternalIdentifierDetail `json:"external-identifier,omitempty"`
	Path               string                     `json:"path,omitempty"`
}

type ExternalIdentifierDetail struct {
	CreatedDate            *Int64Value  `json:"created-date,omitempty"`
	LastModifiedDate       *Int64Value  `json:"last-modified-date,omitempty"`
	Source                 *Source      `json:"source,omitempty"` // Note que o SourceClientID aqui é um objeto OrcidIdentifierInfo
	ExternalIDType         string       `json:"external-id-type,omitempty"`
	ExternalIDValue        string       `json:"external-id-value,omitempty"`
	ExternalIDURL          *StringValue `json:"external-id-url,omitempty"`
	ExternalIDRelationship string       `json:"external-id-relationship,omitempty"`
	Visibility             string       `json:"visibility,omitempty"`
	Path                   string       `json:"path,omitempty"`
	PutCode                int64        `json:"put-code"`
	DisplayIndex           int          `json:"display-index"`
}

type ActivitiesSummary struct {
	LastModifiedDate  *Int64Value         `json:"last-modified-date,omitempty"`
	Distinctions      *AffiliationSection `json:"distinctions,omitempty"`
	Educations        *AffiliationSection `json:"educations,omitempty"`
	Employments       *AffiliationSection `json:"employments,omitempty"`
	Fundings          *FundingSection     `json:"fundings,omitempty"`
	InvitedPositions  *AffiliationSection `json:"invited-positions,omitempty"`
	Memberships       *AffiliationSection `json:"memberships,omitempty"`
	PeerReviews       *PeerReviewSection  `json:"peer-reviews,omitempty"`
	Qualifications    *AffiliationSection `json:"qualifications,omitempty"`
	ResearchResources *GroupSection       `json:"research-resources,omitempty"` // Estrutura interna do grupo não detalhada
	Services          *AffiliationSection `json:"services,omitempty"`
	Works             *WorkSection        `json:"works,omitempty"`
	Path              string              `json:"path,omitempty"`
}

type AffiliationSection struct {
	LastModifiedDate *Int64Value        `json:"last-modified-date,omitempty"`
	AffiliationGroup []AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string             `json:"path,omitempty"`
}

type AffiliationGroup struct {
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	Summaries        []Summary        `json:"summaries,omitempty"`
}

type Summary struct {
	DistinctionSummary   *DistinctionSummary   `json:"distinction-summary,omitempty"`
	EducationSummary     *EducationSummary     `json:"education-summary,omitempty"`
	EmploymentSummary    *EmploymentSummary    `json:"employment-summary,omitempty"`
	QualificationSummary *QualificationSummary `json:"qualification-summary,omitempty"`
	ServiceSummary       *ServiceSummary       `json:"service-summary,omitempty"`
	// Adicionar outros tipos de sumários aqui se necessário (InvitedPositionSummary, MembershipSummary)
}

type DistinctionSummary struct {
	CreatedDate      *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	PutCode          int64            `json:"put-code"`
	DepartmentName   *string          `json:"department-name,omitempty"`
	RoleTitle        string           `json:"role-title,omitempty"`
	StartDate        *PartialDate     `json:"start-date,omitempty"`
	EndDate          *PartialDate     `json:"end-date,omitempty"`
	Organization     *Organization    `json:"organization,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
	Path             string           `json:"path,omitempty"`
}

type EducationSummary struct {
	CreatedDate      *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	PutCode          int64            `json:"put-code"`
	DepartmentName   *string          `json:"department-name,omitempty"`
	RoleTitle        string           `json:"role-title,omitempty"`
	StartDate        *PartialDate     `json:"start-date,omitempty"`
	EndDate          *PartialDate     `json:"end-date,omitempty"`
	Organization     *Organization    `json:"organization,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
	Path             string           `json:"path,omitempty"`
}

type EmploymentSummary struct {
	CreatedDate      *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	PutCode          int64            `json:"put-code"`
	DepartmentName   *string          `json:"department-name,omitempty"`
	RoleTitle        string           `json:"role-title,omitempty"`
	StartDate        *PartialDate     `json:"start-date,omitempty"`
	EndDate          *PartialDate     `json:"end-date,omitempty"`
	Organization     *Organization    `json:"organization,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
	Path             string           `json:"path,omitempty"`
}

type QualificationSummary struct {
	CreatedDate      *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	PutCode          int64            `json:"put-code"`
	DepartmentName   *string          `json:"department-name,omitempty"`
	RoleTitle        string           `json:"role-title,omitempty"`
	StartDate        *PartialDate     `json:"start-date,omitempty"`
	EndDate          *PartialDate     `json:"end-date,omitempty"`
	Organization     *Organization    `json:"organization,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
	Path             string           `json:"path,omitempty"`
}

type ServiceSummary struct {
	CreatedDate      *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	PutCode          int64            `json:"put-code"`
	DepartmentName   *string          `json:"department-name,omitempty"`
	RoleTitle        string           `json:"role-title,omitempty"`
	StartDate        *PartialDate     `json:"start-date,omitempty"`
	EndDate          *PartialDate     `json:"end-date,omitempty"`
	Organization     *Organization    `json:"organization,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
	Path             string           `json:"path,omitempty"`
}

type FundingSection struct {
	LastModifiedDate *Int64Value    `json:"last-modified-date,omitempty"`
	Group            []FundingGroup `json:"group,omitempty"`
	Path             string         `json:"path,omitempty"`
}

type FundingGroup struct {
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	FundingSummary   []FundingSummary `json:"funding-summary,omitempty"`
}

type FundingSummary struct {
	CreatedDate      *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	Title            *FundingTitle    `json:"title,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	Type             string           `json:"type,omitempty"`
	StartDate        *PartialDate     `json:"start-date,omitempty"`
	EndDate          *PartialDate     `json:"end-date,omitempty"`
	Organization     *Organization    `json:"organization,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
	PutCode          int64            `json:"put-code"`
	Path             string           `json:"path,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty"`
}

type FundingTitle struct {
	Title           *StringValue `json:"title,omitempty"`
	TranslatedTitle *StringValue `json:"translated-title,omitempty"`
}

type PeerReviewSection struct {
	LastModifiedDate *Int64Value       `json:"last-modified-date,omitempty"`
	Group            []PeerReviewGroup `json:"group,omitempty"`
	Path             string            `json:"path,omitempty"`
}

type PeerReviewGroup struct {
	LastModifiedDate *Int64Value          `json:"last-modified-date,omitempty"`
	ExternalIDs      *ExternalIDsList     `json:"external-ids,omitempty"`
	PeerReviewGroup  []PeerReviewSubGroup `json:"peer-review-group,omitempty"` // Nome do campo no JSON é "peer-review-group"
}

type PeerReviewSubGroup struct { // Renomeado para evitar conflito com PeerReviewGroup (pai)
	LastModifiedDate  *Int64Value         `json:"last-modified-date,omitempty"`
	ExternalIDs       *ExternalIDsList    `json:"external-ids,omitempty"`
	PeerReviewSummary []PeerReviewSummary `json:"peer-review-summary,omitempty"`
}

type PeerReviewSummary struct {
	CreatedDate           *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate      *Int64Value      `json:"last-modified-date,omitempty"`
	Source                *Source          `json:"source,omitempty"`
	ReviewerRole          string           `json:"reviewer-role,omitempty"`
	ExternalIDs           *ExternalIDsList `json:"external-ids,omitempty"` // Note: Pode ter estrutura específica
	ReviewURL             *StringValue     `json:"review-url,omitempty"`
	ReviewType            string           `json:"review-type,omitempty"`
	CompletionDate        *PartialDate     `json:"completion-date,omitempty"`
	ReviewGroupID         string           `json:"review-group-id,omitempty"` // ISSN
	ConveningOrganization *Organization    `json:"convening-organization,omitempty"`
	Visibility            string           `json:"visibility,omitempty"`
	PutCode               int64            `json:"put-code"`
	Path                  string           `json:"path,omitempty"`
	DisplayIndex          string           `json:"display-index,omitempty"`
}

type GroupSection struct { // Para seções como research-resources
	LastModifiedDate *Int64Value   `json:"last-modified-date,omitempty"`
	Group            []interface{} `json:"group,omitempty"` // Grupo genérico, pois a estrutura interna não foi detalhada
	Path             string        `json:"path,omitempty"`
}

type WorkSection struct {
	LastModifiedDate *Int64Value `json:"last-modified-date,omitempty"`
	Group            []WorkGroup `json:"group,omitempty"`
	Path             string      `json:"path,omitempty"`
}

type WorkGroup struct {
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	WorkSummary      []WorkSummary    `json:"work-summary,omitempty"`
}

type WorkSummary struct {
	PutCode          int64            `json:"put-code"`
	CreatedDate      *Int64Value      `json:"created-date,omitempty"`
	LastModifiedDate *Int64Value      `json:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	Title            *WorkTitle       `json:"title,omitempty"`
	ExternalIDs      *ExternalIDsList `json:"external-ids,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	Type             string           `json:"type,omitempty"`
	PublicationDate  *PartialDate     `json:"publication-date,omitempty"`
	JournalTitle     *StringValue     `json:"journal-title,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
	Path             string           `json:"path,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty"`
}

type WorkTitle struct {
	Title           *StringValue `json:"title,omitempty"`
	Subtitle        *StringValue `json:"subtitle,omitempty"`
	TranslatedTitle *StringValue `json:"translated-title,omitempty"`
}

// Estruturas Comuns Aninhadas e Detalhes de ExternalID
type Source struct {
	SourceOrcid             *OrcidIdentifierInfo `json:"source-orcid,omitempty"`
	SourceClientID          *ClientID            `json:"source-client-id,omitempty"` // Pode ser null ou objeto ClientID
	SourceName              *StringValue         `json:"source-name,omitempty"`
	AssertionOriginOrcid    *OrcidIdentifierInfo `json:"assertion-origin-orcid,omitempty"`
	AssertionOriginClientID *ClientID            `json:"assertion-origin-client-id,omitempty"` // Pode ser null ou objeto ClientID
	AssertionOriginName     *StringValue         `json:"assertion-origin-name,omitempty"`
}

type ClientID struct { // Usado em Source quando source-client-id é um objeto
	URI  string `json:"uri,omitempty"`
	Path string `json:"path,omitempty"`
	Host string `json:"host,omitempty"`
}

type PartialDate struct {
	Year  *StringValue `json:"year,omitempty"`
	Month *StringValue `json:"month,omitempty"`
	Day   *StringValue `json:"day,omitempty"`
}

type Organization struct {
	Name                      string                     `json:"name,omitempty"`
	Address                   *OrganizationAddress       `json:"address,omitempty"`
	DisambiguatedOrganization *DisambiguatedOrganization `json:"disambiguated-organization,omitempty"`
}

type OrganizationAddress struct {
	City    string  `json:"city,omitempty"`
	Region  *string `json:"region,omitempty"`
	Country string  `json:"country,omitempty"`
}

type DisambiguatedOrganization struct {
	DisambiguatedOrganizationIdentifier string `json:"disambiguated-organization-identifier,omitempty"`
	DisambiguationSource                string `json:"disambiguation-source,omitempty"`
}

type ExternalIDsList struct { // Usado em AffiliationGroup, FundingSummary, WorkSummary, etc.
	ExternalID []ExternalID `json:"external-id,omitempty"`
}

type ExternalID struct {
	ExternalIDType         string        `json:"external-id-type,omitempty"`
	ExternalIDValue        string        `json:"external-id-value,omitempty"`
	ExternalIDNormalized   *NormalizedID `json:"external-id-normalized,omitempty"`   // Pode ser null
	ExternalIDURLEmbedded  *StringValue  `json:"external-id-url,omitempty"`          // Nome do campo no JSON é "external-id-url"
	ExternalIDRelationship *string       `json:"external-id-relationship,omitempty"` // Pode ser null ou string
}

type NormalizedID struct {
	Value     string `json:"value,omitempty"`
	Transient bool   `json:"transient"`
}
