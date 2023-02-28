package kredit

import (
	"log"
	"math"
	"strconv"
)

type Service interface {
	GetChecklistPencairan(p string, l string) (ResponseChecklistPencairan, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetChecklistPencairan(p string, l string) (ResponseChecklistPencairan, error) {
	// log.Println("page:", req.Page)
	// log.Println("limit:", req.Limit)
	// req.Sanitize()
	log.Println("page:", p)
	log.Println("limit:", l)

	page, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return ResponseChecklistPencairan{}, err
	}
	limit, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		return ResponseChecklistPencairan{}, err
	}

	fields := []string{"approval_status = ?"}
	values := []any{"9"}

	records, countRecord, err := s.repo.GetChecklistPencairan(int(page), int(limit), fields, values)
	if err != nil {
		return ResponseChecklistPencairan{}, err
	}
	countPage := math.Ceil(float64(countRecord) / float64(limit))

	res := ResponseChecklistPencairan{
		Records:     records,
		CountRecord: countRecord,
		CountPage:   int(countPage),
	}
	return res, nil
}
