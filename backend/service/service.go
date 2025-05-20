package service

import (
	"main/model"
	"main/repository"
)

// service layer here
// only business logic here, no http requests.
// it can call the model and repository layer to get data from the database and/or make requests.

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) SearchBiography(orcid_id string) (model.BiographyData, error) {
	orcid_biography, err := s.repository.GetPersonData(orcid_id)
	if err != nil {
		return model.BiographyData{}, err
	}

	biography_data := s.repository.ProcessOrcidBiography(orcid_biography)

	return biography_data, nil
}
