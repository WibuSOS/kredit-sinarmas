package kredit

import (
	"sinarmas/kredit-sinarmas/models"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	GetChecklistPencairan(page int, limit int, fields []string, values []any) ([]RecordChecklistPencairan, int, error)
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
