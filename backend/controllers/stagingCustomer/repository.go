package stagingCustomer

import (
	"errors"
	"log"
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

	validate(r.db, dirtyCustomerList)

	return dirtyCustomerList, nil
}

func validate(db *gorm.DB, dirtyCustomerList []models.StagingCustomer) {
	for i, customer := range dirtyCustomerList {
		customerDataTest := models.CustomerDataTab{}
		if err := db.Take(&customerDataTest, "ppk = ?", strings.TrimSpace(customer.CustomerPpk)).Error; err == nil {
			dirtyCustomerList[i].ScFlag = "8"
			log.Println("ppk")
			log.Println("ID:", customer.ID, "staging:", customer.CustomerPpk, "customer:", customerDataTest.PPK)
		}

		companyDataTest := models.MstCompanyTab{}
		if err := db.Take(&companyDataTest, "company_short_name = ?", strings.TrimSpace(customer.ScCompany)).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			dirtyCustomerList[i].ScFlag = "8"
			log.Println("company")
			log.Println("ID:", customer.ID, "staging:", customer.ScCompany, "company:", companyDataTest.CompanyShortName)
		}

		branchDataTest := models.BranchTab{}
		if err := db.Take(&branchDataTest, "code = ?", strings.TrimSpace(customer.ScBranchCode)).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			dirtyCustomerList[i].ScFlag = "8"
			log.Println("branch")
			log.Println("ID:", customer.ID, "staging:", customer.ScBranchCode, "branch:", branchDataTest.Code)
		}

		currentTime := time.Now()
		if tglPk, _ := time.Parse("2006-01-02", strings.TrimSpace(customer.LoanTglPk)); tglPk.Year() != currentTime.Year() || tglPk.Month() != currentTime.Month() {
			dirtyCustomerList[i].ScFlag = "8"
			log.Println("tgl_pk")
			log.Println(
				"ID:", customer.ID,
				"staging:", customer.LoanTglPk,
				"stagingYear:", tglPk.Year(),
				"stagingMonth:", tglPk.Month(),
				"currentTime:", currentTime.String(),
				"currentTimeYear:", currentTime.Year(),
				"currentTimeMonth:", currentTime.Month(),
			)
		}

		if strings.TrimSpace(customer.CustomerIdType) == "1" && strings.TrimSpace(customer.CustomerIdNumber) == "" {
			dirtyCustomerList[i].ScFlag = "8"
			log.Println("customer_id")
			log.Println("ID:", customer.ID, "id_type:", customer.CustomerIdType, "id_number:", customer.CustomerIdNumber)
		}
	}
}
