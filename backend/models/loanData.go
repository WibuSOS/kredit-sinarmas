package models

import "time"

type LoanDataTab struct {
	Custcode             string    `json:"custcode" gorm:"unique;not null;type:varchar(25)"`
	Branch               string    `json:"branch" gorm:"type:varchar(50)"`
	OTR                  float64   `json:"otr" gorm:"type:decimal(16,2)"`
	DownPayment          float64   `json:"down_payment" gorm:"type:decimal(16,2)"`
	LoanAmount           float64   `json:"loan_amount" gorm:"type:decimal(16,2)"`
	LoanPeriod           string    `json:"loan_period" gorm:"type:varchar(6)"`
	InterestType         uint8     `json:"interest_type" gorm:"type:smallint"`
	InterestFlat         float32   `json:"interest_flat" gorm:"type:real"`
	InterestEffective    float32   `json:"interest_effective" gorm:"type:real"`
	EffectivePaymentType uint8     `json:"effective_payment_type" gorm:"type:smallint"`
	AdminFee             float64   `json:"admin_fee" gorm:"type:decimal(16,2)"`
	MonthlyPayment       float64   `json:"monthly_payment" gorm:"type:decimal(16,2)"`
	InputDate            time.Time `json:"input_date" gorm:"type:timestamp"`
	LastModified         time.Time `json:"last_modified" gorm:"type:timestamp"`
	ModifiedBy           string    `json:"modified_by" gorm:"type:varchar(20)"`
	Inputdate            time.Time `json:"inputdate" gorm:"column:inputdate;type:timestamp"`
	InputBy              string    `json:"input_by" gorm:"type:varchar(50)"`
	Lastmodified         time.Time `json:"lastmodified" gorm:"column:lastmodified;type:timestamp"`
	Modifiedby           string    `json:"modifiedby" gorm:"column:modifiedby;type:varchar(50)"`
}

func (LoanDataTab) TableName() string {
	return "loan_data_tab"
}
