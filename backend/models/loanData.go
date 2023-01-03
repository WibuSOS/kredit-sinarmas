package models

import (
	"time"

	"gorm.io/gorm"
)

type LoanDataTab struct {
	gorm.Model
	Custcode             string    `json:"custcode" gorm:"not null; type:varchar(25); unique"`
	Branch               string    `json:"branch" gorm:"type:varchar(50)"`
	OTR                  float64   `json:"otr" gorm:"type:money"`
	DownPayment          float64   `json:"down_payment" gorm:"type:money"`
	LoanAmount           float64   `json:"loan_amount" gorm:"type:money"`
	LoanPeriod           string    `json:"loan_period" gorm:"type:varchar(6)"`
	InterestType         int8      `json:"interest_type" gorm:"type:smallint"`
	InterestFlat         float32   `json:"interest_flat" gorm:"type:real"`
	InterestEffective    float32   `json:"interest_effective" gorm:"type:real"`
	EffectivePaymentType int8      `json:"effective_payment_type" gorm:"smallint"`
	AdminFee             float64   `json:"admin_fee" gorm:"type:money"`
	MonthlyPayment       float64   `json:"monthly_payment" gorm:"type:money"`
	InputDate            time.Time `json:"input_date" gorm:"type:timestamp"`
	LastModified         time.Time `json:"last_modified" gorm:"type:timestamp"`
	ModifiedBy           string    `json:"modified_by" gorm:"type:varchar(20)"`
	InputDate2           time.Time `json:"inputdate" gorm:"type:timestamp"`
	InputBy              string    `json:"input_by" gorm:"type:varchar(50)"`
	LastModified2        time.Time `json:"lastmodified" gorm:"type:timestamp"`
	ModifiedBy2          string    `json:"modifiedby" gorm:"type:varchar(50)"`
}

func (m *LoanDataTab) TableName() string {
	return "loan_data_tab"
}
