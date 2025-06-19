package service

import (
	"fmt"
	"main/model"
	"main/repository"
	"strconv"
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

func (s *Service) SearchWork(orcid_id string) (model.WorkData, error) {
	orcid_work, err := s.repository.GetWorkData(orcid_id)
	if err != nil {
		return model.WorkData{}, err
	}

	work_data := s.repository.ProcessOrcidWork(orcid_work)

	return work_data, nil
}

func (s *Service) SearchRecord(orcid_id string) (model.RecordDataResponse, error) {
	orcidRecord, err := s.repository.GetRecordData(orcid_id)
	if err != nil {
		return model.RecordDataResponse{}, err
	}

	responseData := s.repository.ProcessOrcidRecord(orcidRecord)

	return responseData, nil
}

func (s *Service) SearchResearchersByName(name string) (model.ResearcherResults, error) {
	ResearcherResults, err := s.repository.SearchByName(name)
	if err != nil {
		return model.ResearcherResults{}, err
	}

	return ResearcherResults, nil
}

func (s * Service) CalculateHIndex(orcid_id string) error {
	// getting all works from an orcid id
	orcid_work, err := s.repository.GetWorkData(orcid_id)
	if err != nil {
		return err
	}

	work_data := s.repository.ProcessOrcidWork(orcid_work) // parsing the data,, we only need the doi 

	var citations []int
	var notFound int = 0

	for _, work := range work_data.Publications {
		count, err := s.repository.GetCitationCount(work.Doi)
		if err != nil {
			fmt.Println("warning: doi not found: ", work.Doi)
			notFound++
		}

		citations = append(citations, count)
	}

	hIndex := s.repository.CalculateHIndex(citations)

	fmt.Println("\nh index for author {" + orcid_id + "}: " + strconv.Itoa(hIndex))
	fmt.Println("n of dois not found: " + strconv.Itoa(notFound))

	return nil
}

func (s* Service ) GetCitationCound(doi string) (int, error) {
	count, err := s.repository.GetCitationCount(doi)
	if err != nil {
		fmt.Println("could not get citation count for doi " + doi + ". Err: " + err.Error())
		return -1, err
	}

	return count, nil
}