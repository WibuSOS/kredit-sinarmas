package kredit

import (
	"errors"
	"strings"
)

type RequestChecklistPencairan struct {
	Page  string `json:"page"`
	Limit string `json:"limit"`
}

func (req *RequestChecklistPencairan) Sanitize() {
	req.Page = strings.TrimSpace(req.Page)
	req.Limit = strings.TrimSpace(req.Limit)

	if req.Page == "" {
		req.Page = "1"
	}

	if req.Limit == "" {
		req.Limit = "10"
	}
}

type RequestUpdateChecklistPencairan struct {
	Custcodes []string `json:"custcodes"`
}

func (req *RequestUpdateChecklistPencairan) Validate() error {
	if len(req.Custcodes) == 0 {
		return errors.New("No Record(s) Submitted")
	}

	return nil
}
