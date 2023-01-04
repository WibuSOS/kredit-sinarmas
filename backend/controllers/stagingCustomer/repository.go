package stagingCustomer

import (
	"sinarmas/kredit-sinarmas/models"
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
