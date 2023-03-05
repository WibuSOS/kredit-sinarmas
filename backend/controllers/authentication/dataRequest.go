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

type RequestChangePassword struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (req RequestChangePassword) Validation() error {
	return nil
}
