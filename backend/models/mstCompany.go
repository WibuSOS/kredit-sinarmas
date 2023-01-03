package models

type MstCompanyTab struct {
	ID               uint    `json:"id" gorm:"type:bigint"`
	CompanyCode      string  `json:"company_code" gorm:"type:varchar(5)"`
	CompanyShortName string  `json:"company_short_name" gorm:"type:varchar(50)"`
	CompanyName      string  `json:"company_name" gorm:"type:varchar(200)"`
	CompanyAddress1  string  `json:"company_address1" gorm:"type:varchar(200)"`
	CompanyAddress2  string  `json:"company_address2" gorm:"type:varchar(200)"`
	CompanyCity      string  `json:"company_city" gorm:"type:varchar(100)"`
	CompanyPhone     string  `json:"company_phone" gorm:"type:varchar(50)"`
	CompanyFax       string  `json:"company_fax" gorm:"type:varchar(50)"`
	BungaEffMin      float32 `json:"bunga_eff_min" gorm:"type:real"`
	BungaEffMax      float32 `json:"bunga_eff_max" gorm:"type:real"`
	BungaFlatMin     float32 `json:"bunga_flat_min" gorm:"type:real"`
	BungaFlatMax     float32 `json:"bunga_flat_max" gorm:"type:real"`
	LAMin            float64 `json:"la_min" gorm:"column:la_min;type:money"`
	LAMax            float64 `json:"la_max" gorm:"column:la_max;type:money"`
	PeriodeMin       string  `json:"periode_min" gorm:"type:varchar(10)"`
	PeriodeMax       string  `json:"periode_max" gorm:"type:varchar(10)"`
}

func (MstCompanyTab) TableName() string {
	return "mst_company_tab"
}
