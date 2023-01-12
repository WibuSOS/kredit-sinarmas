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

	errDescs := validate(r.db, dirtyCustomerList)
	insert(r.db, dirtyCustomerList, errDescs)

	return dirtyCustomerList, nil
}

func validate(db *gorm.DB, dirtyCustomerList []models.StagingCustomer) map[int]string {
	errDescs := map[int]string{}

	log.Println("VALIDATE")
	for i, customer := range dirtyCustomerList {
		customerDataTest := models.CustomerDataTab{}
		if err := db.Take(&customerDataTest, "ppk = ?", strings.TrimSpace(customer.CustomerPpk)).Error; err == nil {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",tidak boleh duplikasi pada PPK"
			} else {
				errDescs[i] = "tidak boleh duplikasi pada PPK"
			}
			log.Println("ppk")
			log.Println("ID:", customer.ID, "staging:", customer.CustomerPpk, "customer:", customerDataTest.PPK)
		}

		companyDataTest := models.MstCompanyTab{}
		if err := db.Take(&companyDataTest, "company_short_name = ?", strings.TrimSpace(customer.ScCompany)).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",company belum terdaftar di master table"
			} else {
				errDescs[i] = "company belum terdaftar di master table"
			}
			log.Println("company")
			log.Println("ID:", customer.ID, "staging:", customer.ScCompany, "company:", companyDataTest.CompanyShortName)
		}

		branchDataTest := models.BranchTab{}
		if err := db.Take(&branchDataTest, "code = ?", strings.TrimSpace(customer.ScBranchCode)).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",branch belum terdaftar di master table"
			} else {
				errDescs[i] = "branch belum terdaftar di master table"
			}
			log.Println("branch")
			log.Println("ID:", customer.ID, "staging:", customer.ScBranchCode, "branch:", branchDataTest.Code)
		}

		currentTime := time.Now()
		if tglPk, _ := time.Parse("2006-01-02", strings.TrimSpace(customer.LoanTglPk)); tglPk.Year() != currentTime.Year() || tglPk.Month() != currentTime.Month() {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",tanggal PK harus sama dengan tahun dan bulan sekarang"
			} else {
				errDescs[i] = "tanggal PK harus sama dengan tahun dan bulan sekarang"
			}
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
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",nomor ID harus terisi"
			} else {
				errDescs[i] = "nomor ID harus terisi"
			}
			log.Println("customer_id")
			log.Println("ID:", customer.ID, "id_type:", customer.CustomerIdType, "id_number:", customer.CustomerIdNumber)
		}

		// regex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]+`)
		// regex.MatchString(strings.TrimSpace(customer.CustomerName))
		if matched := strings.ContainsAny(strings.TrimSpace(customer.CustomerName), `!@#$%^&*()_+-=[]{};':"\|,.<>/?~`); strings.TrimSpace(customer.CustomerName) == "" || matched {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",nama debitur tidak boleh mengandung karakter spesial"
			} else {
				errDescs[i] = "nama debitur tidak boleh mengandung karakter spesial"
			}
			log.Println("nama_debitur")
			log.Println("ID:", customer.ID, "customer_name:", customer.CustomerName)
		}

		if strings.TrimSpace(customer.VehicleBpkb) == "" {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",BPKB harus terisi"
			} else {
				errDescs[i] = "BPKB harus terisi"
			}
			log.Println("vehicle_bpkb")
			log.Println("ID:", customer.ID, "vehicle_bpkb:", customer.VehicleBpkb)
		}

		if strings.TrimSpace(customer.VehicleStnk) == "" {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",STNK harus terisi"
			} else {
				errDescs[i] = "STNK harus terisi"
			}
			log.Println("vehicle_stnk")
			log.Println("ID:", customer.ID, "vehicle_stnk:", customer.VehicleStnk)
		}

		if strings.TrimSpace(customer.VehicleEngineNo) == "" {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",nomor mesin harus terisi"
			} else {
				errDescs[i] = "nomor mesin harus terisi"
			}
			log.Println("vehicle_engine_no")
			log.Println("ID:", customer.ID, "staging_vehicle_engine_no:", customer.VehicleEngineNo)
		}

		vehicleEngineNoDataTest := models.VehicleDataTab{}
		if err := db.Take(&vehicleEngineNoDataTest, "engine_no = ?", strings.TrimSpace(customer.VehicleEngineNo)).Error; err == nil {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",tidak boleh duplikasi pada nomor mesin"
			} else {
				errDescs[i] = "tidak boleh duplikasi pada nomor mesin"
			}
			log.Println("vehicle_engine_no")
			log.Println("ID:", customer.ID, "staging_vehicle_engine_no:", customer.VehicleEngineNo, "vehicle_engine_no:", vehicleEngineNoDataTest.EngineNo)
		}

		if strings.TrimSpace(customer.VehicleChasisNo) == "" {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",nomor chasis harus terisi"
			} else {
				errDescs[i] = "nomor chasis harus terisi"
			}
			log.Println("vehicle_chasis_no")
			log.Println("ID:", customer.ID, "staging_vehicle_chasis_no:", customer.VehicleChasisNo)
		}

		vehicleChasisNoDataTest := models.VehicleDataTab{}
		if err := db.Take(&vehicleChasisNoDataTest, "chasis_no = ?", strings.TrimSpace(customer.VehicleChasisNo)).Error; err == nil {
			dirtyCustomerList[i].ScFlag = "8"
			if _, exists := errDescs[i]; exists {
				errDescs[i] = errDescs[i] + ",tidak boleh duplikasi pada nomor chasis"
			} else {
				errDescs[i] = "tidak boleh duplikasi pada nomor chasis"
			}
			log.Println("vehicle_chasis_no")
			log.Println("ID:", customer.ID, "staging_vehicle_chasis_no:", customer.VehicleChasisNo, "vehicle_chasis_no:", vehicleChasisNoDataTest.ChasisNo)
		}

		if strings.TrimSpace(dirtyCustomerList[i].ScFlag) == "0" {
			dirtyCustomerList[i].ScFlag = "1"
		}

		log.Printf("%+v", dirtyCustomerList[i])
		log.Println("errDesc:", errDescs[i])
	}

	return errDescs
}

func insert(db *gorm.DB, dirtyCustomerList []models.StagingCustomer, errDescs map[int]string) {
	log.Println("INSERT")
	cleanCustomers := []models.CustomerDataTab{}
	cleanLoans := []models.LoanDataTab{}
	cleanVehicles := []models.VehicleDataTab{}
	cleanErrors := []models.StagingError{}
	currentTime := time.Now()

	for i, dirtyCustomer := range dirtyCustomerList {
		log.Println("ID:", dirtyCustomer.ID)
		log.Println("ScFlag:", dirtyCustomer.ScFlag)
		log.Println("errDesc:", errDescs[i])

		if dirtyCustomer.ScFlag == "8" {
			cleanErrors = append(cleanErrors, createStagingError(dirtyCustomer, currentTime, errDescs[i]))
			continue
		}

		cleanCustomers = append(cleanCustomers, createCustomerData(dirtyCustomer, currentTime))
		cleanLoans = append(cleanLoans, createLoanData(dirtyCustomer, currentTime))
		cleanVehicles = append(cleanVehicles, createVehicleData(dirtyCustomer, currentTime))
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&cleanErrors).Error; err != nil {
			return err
		}

		if err := tx.Create(&cleanCustomers).Error; err != nil {
			return err
		}

		if err := tx.Create(&cleanLoans).Error; err != nil {
			return err
		}

		if err := tx.Create(&cleanVehicles).Error; err != nil {
			return err
		}

		if err := tx.Model(&dirtyCustomerList).Update("sc_flag", "lol").Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		log.Println(err.Error())
	}
}

func createCustomerData(dirtyCustomer models.StagingCustomer, currentTime time.Time) models.CustomerDataTab {
	customerData := models.CustomerDataTab{}

	return customerData
}

func createLoanData(dirtyCustomer models.StagingCustomer, currentTime time.Time) models.LoanDataTab {
	loanData := models.LoanDataTab{}

	return loanData
}

func createVehicleData(dirtyCustomer models.StagingCustomer, currentTime time.Time) models.VehicleDataTab {
	vehicleData := models.VehicleDataTab{}

	return vehicleData
}

func createStagingError(dirtyCustomer models.StagingCustomer, currentTime time.Time, errDesc string) models.StagingError {
	stagingError := models.StagingError{}

	return stagingError
}
