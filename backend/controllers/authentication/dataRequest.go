package authentication

import (
	"fmt"
	"strings"
)

type DataRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req DataRequest) Validation() error {
	if strings.TrimSpace(req.Username) == "" {
		return fmt.Errorf("invalid username")
	}

	if req.Password == "" {
		return fmt.Errorf("invalid password")
	}

	return nil
}
