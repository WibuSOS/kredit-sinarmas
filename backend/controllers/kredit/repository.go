package kredit

import (
	"log"
	"sinarmas/kredit-sinarmas/models"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	GetChecklistPencairan(page int, limit int, fields []string, values []any) ([]RecordChecklistPencairan, int, error)
	UpdateChecklistPencairan(fields []string, values []string) error
	GetDrawdownReport(page int, limit int, fields []string, values []any) ([]RecordDrawdownReport, int, []RecordCompany, []RecordBranch, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetChecklistPencairan(page int, limit int, fields []string, values []any) ([]RecordChecklistPencairan, int, error) {
	var records []RecordChecklistPencairan
	offset := (page - 1) * limit

	err := r.db.
		Model(&models.CustomerDataTab{}).
		Limit(limit).
		Offset(offset).
		Select(
			"customer_data_tab.custcode, "+
				"customer_data_tab.ppk, "+
				"customer_data_tab.name, "+
				"customer_data_tab.channeling_company, "+
				"customer_data_tab.drawdown_date, "+
				"loan_data_tab.loan_amount, "+
				"loan_data_tab.loan_period, "+
				"loan_data_tab.interest_effective, "+
				"customer_data_tab.approval_status",
		).
		Joins("INNER JOIN loan_data_tab ON customer_data_tab.custcode = loan_data_tab.custcode").
		Where(strings.Join(fields, " AND "), values...).
		Order("drawdown_date DESC").
		Scan(&records).Error

	if err != nil {
		return []RecordChecklistPencairan{}, 0, err
	}

	countRecord, err := getCountChecklistPencairan(r.db, fields, values)

	if err != nil {
		return []RecordChecklistPencairan{}, 0, err
	}

	return records, countRecord, nil
}

func getCountChecklistPencairan(db *gorm.DB, fields []string, values []any) (int, error) {
	type resultFormat struct {
		Size int
	}
	var result resultFormat

	err := db.
		Model(&models.CustomerDataTab{}).
		Select("COUNT(customer_data_tab.custcode) as size").
		Joins("INNER JOIN loan_data_tab ON customer_data_tab.custcode = loan_data_tab.custcode").
		Where(strings.Join(fields, " AND "), values...).
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return result.Size, nil
}

func (r *repository) UpdateChecklistPencairan(fields []string, values []string) error {
	err := r.db.
		Model(&models.CustomerDataTab{}).
		Where(strings.Join(fields, " AND "), values).
		Update("approval_status", "0").Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetDrawdownReport(page int, limit int, fields []string, values []any) ([]RecordDrawdownReport, int, []RecordCompany, []RecordBranch, error) {
	var records []RecordDrawdownReport
	offset := (page - 1) * limit

	err := r.db.
		Model(&models.CustomerDataTab{}).
		Limit(limit).
		Offset(offset).
		Select(
			"customer_data_tab.custcode, "+
				"customer_data_tab.ppk, "+
				"customer_data_tab.name, "+
				"customer_data_tab.channeling_company, "+
				"customer_data_tab.drawdown_date, "+
				"loan_data_tab.loan_amount, "+
				"loan_data_tab.loan_period, "+
				"loan_data_tab.interest_effective, "+
				"customer_data_tab.approval_status",
		).
		Joins("INNER JOIN loan_data_tab ON customer_data_tab.custcode = loan_data_tab.custcode").
		Where(strings.Join(fields, " AND "), values...).
		Order("drawdown_date DESC").
		Scan(&records).Error

	if err != nil {
		return []RecordDrawdownReport{}, 0, []RecordCompany{}, []RecordBranch{}, err
	}

	countRecord, err := getCountChecklistPencairan(r.db, fields, values)

	if err != nil {
		return []RecordDrawdownReport{}, 0, []RecordCompany{}, []RecordBranch{}, err
	}

	companies, err := getCompany(r.db)
	if err != nil {
		return []RecordDrawdownReport{}, 0, []RecordCompany{}, []RecordBranch{}, err
	}

	branches, err := getBranch(r.db)
	if err != nil {
		return []RecordDrawdownReport{}, 0, []RecordCompany{}, []RecordBranch{}, err
	}

	return records, countRecord, companies, branches, nil
}

func getCompany(db *gorm.DB) ([]RecordCompany, error) {
	var companies []RecordCompany

	if err := db.Model(&models.MstCompanyTab{}).Select("company_code, company_short_name").Scan(&companies).Error; err != nil {
		log.Println("masuk getCompany")
		return []RecordCompany{}, err
	}

	return companies, nil
}

func getBranch(db *gorm.DB) ([]RecordBranch, error) {
	var branches []RecordBranch

	if err := db.Model(&models.BranchTab{}).Select("code, description").Order("code").Scan(&branches).Error; err != nil {
		return []RecordBranch{}, err
	}

	return branches, nil
}
