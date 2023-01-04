package models

import "time"

type VehicleDataTab struct {
	Custcode       string    `json:"custcode" gorm:"not null;type:varchar(25)"`
	Brand          uint      `json:"brand"`
	Type           string    `json:"type" gorm:"type:varchar(100)"`
	Year           string    `json:"year" gorm:"type:varchar(4)"`
	Golongan       uint8     `json:"golongan" gorm:"type:smallint"`
	Jenis          string    `json:"jenis" gorm:"type:varchar(200)"`
	Status         uint8     `json:"status" gorm:"type:smallint"`
	Color          string    `json:"color" gorm:"type:varchar(20)"`
	PoliceNo       string    `json:"police_no" gorm:"type:varchar(20)"`
	EngineNo       string    `json:"engine_no" gorm:"type:varchar(20)"`
	ChasisNo       string    `json:"chasis_no" gorm:"type:varchar(20)"`
	BPKB           string    `json:"bpkb" gorm:"type:varchar(20)"`
	RegisterNo     string    `json:"register_no" gorm:"type:varchar(50)"`
	STNK           string    `json:"stnk" gorm:"type:varchar(50)"`
	StnkAddress1   string    `json:"stnk_address1" gorm:"type:varchar(40)"`
	StnkAddress2   string    `json:"stnk_address2" gorm:"type:varchar(40)"`
	StnkCity       string    `json:"stnk_city" gorm:"type:varchar(20)"`
	DealerID       uint      `json:"dealer_id"`
	Inputdate      time.Time `json:"inputdate" gorm:"column:inputdate;type:timestamp"`
	Inputby        string    `json:"inputby" gorm:"column:inputby;type:varchar(50)"`
	Lastmodified   time.Time `json:"lastmodified" gorm:"column:lastmodified;type:timestamp"`
	Modifiedby     string    `json:"modifiedby" gorm:"column:modifiedby;type:varchar(50)"`
	TglStnk        time.Time `json:"tgl_stnk" gorm:"type:timestamp"`
	TglBpkb        time.Time `json:"tgl_bpkb" gorm:"type:timestamp"`
	TglPolis       time.Time `json:"tgl_polis" gorm:"type:timestamp"`
	PolisNo        string    `json:"polis_no" gorm:"type:varchar(17)"`
	CollateralID   uint64    `json:"collateral_id" gorm:"type:bigint"`
	Ketagunan      string    `json:"ketagunan" gorm:"type:text"`
	AgunanLbu      string    `json:"agunan_lbu" gorm:"type:varchar(10)"`
	Dealer         string    `json:"dealer" gorm:"type:varchar(100)"`
	AddressDealer1 string    `json:"address_dealer1" gorm:"type:varchar(100)"`
	AddressDealer2 string    `json:"address_dealer2" gorm:"type:varchar(100)"`
	CityDealer     string    `json:"city_dealer" gorm:"type:varchar(100)"`
}

func (VehicleDataTab) TableName() string {
	return "vehicle_data_tab"
}
