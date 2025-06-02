package repository

// all the data fetching and treatment logic here
// its called from the service layer
//

type Repository struct {
	apiLink string
	// relevant fields here, maybe orcid client, etc
}

func NewRepository() *Repository {
	return &Repository{apiLink: "https://pub.orcid.org/v3.0/"}
}
