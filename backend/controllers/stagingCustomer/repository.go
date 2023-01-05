package stagingCustomer

import (
	"errors"
	"sinarmas/kredit-sinarmas/models"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	ValidateAndMigrate() ([]models.StagingCustomer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ValidateAndMigrate() ([]models.StagingCustomer, error) {
	var dirtyCustomerList []models.StagingCustomer
	currentTime := time.Now()

	err := r.db.
		Find(&dirtyCustomerList,
			"sc_flag = ? "+
				"AND DATE_PART('year',sc_create_date) = ? "+
				"AND DATE_PART('month',sc_create_date) = ? "+
				"AND DATE_PART('day',sc_create_date) = ?",
			"0", currentTime.Year(), currentTime.Month(), currentTime.Day()).
		Error
	if err != nil {
		return []models.StagingCustomer{}, err
	}

	return dirtyCustomerList, nil
}

func validate(db *gorm.DB, dirtyCustomerList []models.StagingCustomer) {
	for i, customer := range dirtyCustomerList {
		if err := db.Take(&models.CustomerDataTab{}, "ppk = ?", strings.ToLower(strings.TrimSpace(customer.CustomerPpk))).Error; err == nil {
			dirtyCustomerList[i].ScFlag = "8"
			continue
		}

		if err := db.Take(&models.MstCompanyTab{}, "company_short_name = ?", strings.ToLower(strings.TrimSpace(customer.ScCompany))).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			dirtyCustomerList[i].ScFlag = "8"
			continue
		}
	}
}
