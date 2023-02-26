package kredit

import (
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
