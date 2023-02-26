package kredit

import (
	"math"
	"strconv"
)

type Service interface {
	GetChecklistPencairan(req *RequestChecklistPencairan) (ResponseChecklistPencairan, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetChecklistPencairan(req *RequestChecklistPencairan) (ResponseChecklistPencairan, error) {
	req.Sanitize()

	page, err := strconv.ParseInt(req.Page, 10, 64)
	if err != nil {
		return ResponseChecklistPencairan{}, err
	}
	limit, err := strconv.ParseInt(req.Limit, 10, 64)
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
