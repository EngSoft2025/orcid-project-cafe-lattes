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

func (s *Service) SearchWork(orcid_id string) (model.WorkData, error) {
	orcid_work, err := s.repository.GetWorkData(orcid_id)
	if err != nil {
		return model.WorkData{}, err
	}

	work_data := s.repository.ProcessOrcidWork(orcid_work)

	return work_data, nil
}

func (s *Service) SearchKeyword(orcid_id string) (model.KeywordData, error) {
	orcid_keyword, err := s.repository.GetKeywordData(orcid_id)
	if err != nil {
		return model.KeywordData{}, err
	}

	keyword_data := s.repository.ProcessOrcidKeyword(orcid_keyword)

	return keyword_data, nil
}

func (s *Service) SearchEmployment(orcid_id string) (model.EmploymentData, error) {
	orcid_employment, err := s.repository.GetEmploymentData(orcid_id)
	if err != nil {
		return model.EmploymentData{}, err
	}

	employment_data := s.repository.ProcessOrcidEmployment(orcid_employment)

	return employment_data, nil
}

func (s *Service) SearchReseacherUrls(orcid_id string) (model.ReseacherUrlData, error) {
	orcid_reseacher_url, err := s.repository.GetReseacherData(orcid_id)
	if err != nil {
		return model.ReseacherUrlData{}, err
	}

	reseacher_url_data := s.repository.ProcessOrcidResearcherUrl(orcid_reseacher_url)

	return reseacher_url_data, nil
}

func (s *Service) SearchEducation(orcid_id string) (model.EducationData, error) {
	orcid_education, err := s.repository.GetEducationData(orcid_id)
	if err != nil {
		return model.EducationData{}, err
	}

	education_data := s.repository.ProcessOrcidEducation(orcid_education)

	return education_data, nil
}

func (s *Service) SearchEmail(orcid_id string) (model.EmailData, error) {
	orcid_email, err := s.repository.GetEmailData(orcid_id)
	if err != nil {
		return model.EmailData{}, err
	}

	email_data := s.repository.ProcessOrcidEmail(orcid_email)

	return email_data, nil
}
