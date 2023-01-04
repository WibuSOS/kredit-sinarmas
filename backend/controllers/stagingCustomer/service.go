package stagingCustomer

import "sinarmas/kredit-sinarmas/models"

type Service interface {
	ValidateAndMigrate() ([]models.StagingCustomer, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) ValidateAndMigrate() ([]models.StagingCustomer, error) {
	res, err := s.repo.ValidateAndMigrate()
	if err != nil {
		return []models.StagingCustomer{}, err
	}
	return res, nil
}
