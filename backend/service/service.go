package service

import (
	"fmt"
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

func (s *Service) Search(name, orcid_id string) (string, error) {
	// todo: implement the logic here

	// call repository methods here to fetch from api, treat data, etc

	// service functions only orchestrate the logic , they don't do anything by themselves

	fmt.Println("service doing stuff with payload:", name, orcid_id)
	return "", nil
}
