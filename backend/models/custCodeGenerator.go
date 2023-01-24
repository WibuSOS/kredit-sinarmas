package models

import "gorm.io/gorm"

type IdGeneratorTab struct {
	gorm.Model
	Code  string `gorm:"unique;type:varchar(3)"`
	Value uint   `gorm:"type:bigint"`
	Digit uint   `gorm:"type:bigint"`
}

func (IdGeneratorTab) TableName() string {
	return "id_generator_tab"
}
