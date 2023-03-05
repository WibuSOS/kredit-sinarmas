package kredit

import (
	"log"
	"math"
	"strconv"
)

type Service interface {
	GetChecklistPencairan(p string, l string) (ResponseChecklistPencairan, error)
	UpdateChecklistPencairan(req *RequestUpdateChecklistPencairan) error
	GetDrawdownReport(p string, l string, c string, b string, sD string, eD string, aS string) (ResponseDrawdownReport, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetChecklistPencairan(p string, l string) (ResponseChecklistPencairan, error) {
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

func (s *service) UpdateChecklistPencairan(req *RequestUpdateChecklistPencairan) error {
	if err := req.Validate(); err != nil {
		return err
	}

	fields := []string{"custcode in (?)"}
	values := req.Custcodes

	err := s.repo.UpdateChecklistPencairan(fields, values)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetDrawdownReport(p string, l string, c string, b string, sD string, eD string, aS string) (ResponseDrawdownReport, error) {
	log.Println("page:", p)
	log.Println("limit:", l)
	log.Println("company:", c)
	log.Println("branch:", b)
	log.Println("start_date:", sD)
	log.Println("end_date:", eD)
	log.Println("approval_status:", aS)

	page, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return ResponseDrawdownReport{}, err
	}
	limit, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		return ResponseDrawdownReport{}, err
	}

	fields := []string{}
	values := []any{}

	if c != "" {
		fields = append(fields, "channeling_company = ?")
		values = append(values, c)
	}
	if b != "" {
		fields = append(fields, "branch = ?")
		values = append(values, b)
	}
	if sD != "" {
		fields = append(fields, "drawdown_date >= ?")
		values = append(values, sD)
	}
	if eD != "" {
		fields = append(fields, "drawdown_date <= ?")
		values = append(values, eD)
	}
	if aS == "" {
		fields = append(fields, "approval_status in (?)")
		values = append(values, []string{"0", "1"})
	} else {
		fields = append(fields, "approval_status = ?")
		values = append(values, aS)
	}

	records, countRecord, companies, branches, err := s.repo.GetDrawdownReport(int(page), int(limit), fields, values)
	log.Println(companies, branches)
	if err != nil {
		return ResponseDrawdownReport{}, err
	}
	countPage := math.Ceil(float64(countRecord) / float64(limit))

	res := ResponseDrawdownReport{
		ResponseChecklistPencairan: ResponseChecklistPencairan{
			Records:     records,
			CountRecord: countRecord,
			CountPage:   int(countPage),
		},
		Companies:        companies,
		Branches:         branches,
		ApprovalStatuses: []string{"0", "1"},
	}
	return res, nil
}
