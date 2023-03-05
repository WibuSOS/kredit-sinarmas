package models

import "time"

type SkalaRentalTab struct {
	Custcode      string    `json:"custcode" gorm:"not null;type:varchar(25)"`
	Counter       uint8     `json:"counter" gorm:"type:smallint"`
	OsBalance     float64   `json:"os_balance" gorm:"type:decimal(16,2)"`
	EndBalance    float64   `json:"end_balance" gorm:"type:decimal(16,2)"`
	DueDate       time.Time `json:"due_date" gorm:"type:timestamp"`
	EffRate       float32   `json:"eff_rate" gorm:"type:float"`
	Rental        float64   `json:"rental" gorm:"type:decimal(16,2)"`
	Principle     float64   `json:"principle" gorm:"type:decimal(16,2)"`
	Interest      float64   `json:"interest" gorm:"type:decimal(16,2)"`
	Inputdate     time.Time `json:"inputdate" gorm:"column:inputdate;type:timestamp"`
	Inputby       string    `json:"inputby" gorm:"column:inputby;type:varchar(50)"`
	Lastmodified  time.Time `json:"lastmodified" gorm:"column:lastmodified;type:timestamp"`
	Modifiedby    string    `json:"modifiedby" gorm:"column:modifiedby;type:varchar(50)"`
	PaymentDate   time.Time `json:"payment_date" gorm:"type:timestamp"`
	Penalty       float64   `json:"penalty" gorm:"type:decimal(16,2)"`
	PaymentAmount float64   `json:"payment_amount" gorm:"type:decimal(16,2)"`
	PaymentType   uint8     `json:"payment_type" gorm:"type:smallint"`
}

func (SkalaRentalTab) TableName() string {
	return "skala_rental_tab"
}
