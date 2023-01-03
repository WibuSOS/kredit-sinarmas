package models

import "time"

type VehicleDataTab struct {
	Custcode       string    `json:"custcode" gorm:"not null; type:varchar(25); unique"`
	Brand          int       `json:"brand"`
	Type           string    `json:"type" gorm:"type:varchar(100)"`
	Year           string    `json:"year" gorm:"type:varchar(4)"`
	Golongan       int8      `json:"golongan" gorm:"type:smallint"`
	Jenis          string    `json:"jenis"  gorm:"type:varchar(200)"`
	Status         int8      `json:"status"  gorm:"type:smallint"`
	Color          string    `json:"color"  gorm:"type:varchar(20)"`
	PoliceNo       string    `json:"police_no"  gorm:"type:varchar(20)"`
	EngineNo       string    `json:"engine_no"  gorm:"type:varchar(20)"`
	ChasisNo       string    `json:"chasis_no"  gorm:"type:varchar(20)"`
	Bpkb           string    `json:"bpkb"  gorm:"type:varchar(20)"`
	RegisterNo     string    `json:"register_no"  gorm:"type:varchar(50)"`
	Stnk           string    `json:"stnk"  gorm:"type:varchar(50)"`
	StnkAddress1   string    `json:"stnk_address1"  gorm:"type:varchar(40)"`
	StnkAddress2   string    `json:"stnk_address2" gorm:"type:varchar(40)"`
	StnkCity       string    `json:"stnk_city" gorm:"type:varchar(20)"`
	DealerID       int       `json:"dealer_id"`
	Inputdate      time.Time `json:"inputdate" gorm:"type:timestamp"`
	Inputby        string    `json:"inputby" gorm:"type:varchar(50)"`
	Lastmodified   time.Time `json:"lastmodified" gorm:"type:timestamp"`
	Modifiedby     string    `json:"modifiedby" gorm:"type:varchar(50)"`
	TglStnk        time.Time `json:"tgl_stnk" gorm:"type:timestamp"`
	TglBpkb        time.Time `json:"tgl_bpkb" gorm:"type:timestamp"`
	TglPolis       time.Time `json:"tgl_polis" gorm:"type:timestamp"`
	PolisNo        string    `json:"polis_no" gorm:"type:varchar(17)"`
	CollateralID   int64     `json:"collateral_id" gorm:"type:bigint"`
	Ketagunan      string    `json:"ketagunan" gorm:"type:text"`
	AgunanLbu      string    `json:"agunan_lbu" gorm:"type:varchar(10)"`
	Dealer         string    `json:"dealer" gorm:"type:varchar(100)"`
	AddressDealer1 string    `json:"address_dealer1" gorm:"type:varchar(100)"`
	AddressDealer2 string    `json:"addres_dealer2" gorm:"type:varchar(100)"`
	CityDealer     string    `json:"city_dealer" gorm:"type:varchar(100)"`
}

func (m *VehicleDataTab) TableName() string {
	return "vehicle_data_tab"
}
