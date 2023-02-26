package kredit

import "time"

type RecordChecklistPencairan struct {
	Custcode          string    `json:"custcode"`
	PPK               string    `json:"ppk"`
	Name              string    `json:"name"`
	ChannelingCompany string    `json:"channeling_company"`
	DrawdownDate      time.Time `json:"drawdown_date"`
	LoanAmount        float64   `json:"loan_amount"`
	LoanPeriod        string    `json:"loan_period"`
	InterestEffective float32   `json:"interest_effective"`
	ApprovalStatus    string    `json:"approval_status"`
}
