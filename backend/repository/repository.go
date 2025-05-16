package repository

// all the data fetching and treatment logic here
// its called from the service layer

type Repository struct {
	// relevant fields here, maybe orcid client, etc
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FetchData(name, orcid_id string) (string, error) {
	// todo: implement the logic here

	// make api requests here

	// treat data here

	// return the data here
	return "", nil
}
