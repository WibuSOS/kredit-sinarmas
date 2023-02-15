package automatedService

import (
	"log"
)

type Service interface {
	ValidateAndMigrate()
	GenerateSkalaAngsuran()
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) ValidateAndMigrate() {
	err := s.repo.ValidateAndMigrate()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("ValidateAndMigrate completed")
}

func (s *service) GenerateSkalaAngsuran() {
	err := s.repo.GenerateSkalaAngsuran()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("GenerateSkalaAngsuran completed")
}
