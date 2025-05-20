package service

import (
	"log"
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

func (s *Service) SearchBiography(orcid_id string) model.BiographyData {
	orcid_biography, err := s.repository.GetPersonData(orcid_id)
	if err != nil {
		log.Panicln("Could not get" + orcid_id + " biography data")
	}

	biography_data, err := s.repository.ProcessOrcidBiography(orcid_biography)
	if err != nil {
		log.Println("Could not process" + orcid_id + " biography data")
		log.Panicln("Internal Server Error")
	}

	return biography_data
}
