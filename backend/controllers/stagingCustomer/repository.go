package stagingCustomer

import (
	"errors"
	"fmt"
	"log"
	"sinarmas/kredit-sinarmas/models"
	"strconv"
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
		if tglPk, err := time.Parse("2006-01-02", strings.TrimSpace(customer.LoanTglPk)); err != nil || tglPk.Year() != currentTime.Year() || tglPk.Month() != currentTime.Month() {
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

		newCustCode := generateCustCode(db, dirtyCustomer, currentTime)
		cleanCustomers = append(cleanCustomers, createCustomerData(dirtyCustomer, currentTime, newCustCode))
		cleanLoans = append(cleanLoans, createLoanData(dirtyCustomer, currentTime, newCustCode))
		cleanVehicles = append(cleanVehicles, createVehicleData(dirtyCustomer, currentTime, newCustCode))
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

func generateCustCode(db *gorm.DB, dirtyCustomer models.StagingCustomer, currentTime time.Time) string {
	// lastCustCodeTab := models.IdGeneratorTab{}
	// newCustCodeTab := models.IdGeneratorTab{}
	custCodeTab := models.IdGeneratorTab{}
	mstCompany := models.MstCompanyTab{}
	const appCustCode = "006"

	db.Take(&mstCompany, "company_short_name = ?", strings.TrimSpace(dirtyCustomer.ScCompany))
	companyCode := strings.TrimSpace(mstCompany.CompanyCode)

	if err := db.Last(&custCodeTab, "code = ?", appCustCode).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("last cust code record not found. Initializing from 1...")
		custCodeTab.Code = appCustCode
		custCodeTab.Value = 1
		custCodeTab.Digit = uint(len(fmt.Sprintf("%d", 1)))
		db.Create(&custCodeTab)
	}
	appCustCodeSeqNew := "0000000000" + fmt.Sprintf("%d", custCodeTab.Value)

	month := int(currentTime.Month())
	monthString := ""
	year := currentTime.Year()
	yearString := fmt.Sprintf("%d", year)
	if month < 10 {
		monthString = fmt.Sprintf("0%d", month)
	} else {
		monthString = fmt.Sprintf("%d", month)
	}

	newCustCode := appCustCode + companyCode + yearString + monthString + appCustCodeSeqNew
	custCodeTab.Value += 1
	custCodeTab.Digit = uint(len(fmt.Sprintf("%d", custCodeTab.Value)))
	db.Model(&custCodeTab).Select("Value", "Digit").Updates(&custCodeTab)

	return newCustCode
}

func createCustomerData(dirtyCustomer models.StagingCustomer, currentTime time.Time, custCode string) models.CustomerDataTab {
	birthDate, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(dirtyCustomer.CustomerBirthDate))
	if err != nil {
		log.Println(err.Error())
	}
	idType, err := strconv.ParseUint(strings.TrimSpace(dirtyCustomer.CustomerIdType), 10, 8)
	if err != nil {
		log.Println(err.Error())
	}
	drawdownDate, err := time.Parse("2006-01-02", strings.TrimSpace(dirtyCustomer.LoanTglPk))
	if err != nil {
		log.Println(err.Error())
	}
	tglPkChanneling, err := time.Parse("2006-01-02", strings.TrimSpace(dirtyCustomer.LoanTglPkChanneling))
	if err != nil {
		log.Println(err.Error())
	}

	customerData := models.CustomerDataTab{
		Custcode:          strings.TrimSpace(custCode),
		PPK:               strings.TrimSpace(dirtyCustomer.CustomerPpk),
		Name:              strings.TrimSpace(dirtyCustomer.CustomerName),
		Address1:          strings.TrimSpace(dirtyCustomer.CustomerAddress1),
		Address2:          strings.TrimSpace(dirtyCustomer.CustomerAddress2),
		City:              strings.TrimSpace(dirtyCustomer.CustomerCity),
		Zip:               strings.TrimSpace(dirtyCustomer.CustomerZip),
		BirthPlace:        strings.TrimSpace(dirtyCustomer.CustomerBirthPlace),
		BirthDate:         birthDate,
		IdType:            uint8(idType),
		IdNumber:          strings.TrimSpace(dirtyCustomer.CustomerIdNumber),
		MobileNo:          strings.TrimSpace(dirtyCustomer.CustomerMobileNo),
		DrawdownDate:      drawdownDate,
		TglPkChanneling:   tglPkChanneling,
		MotherMaidenName:  strings.TrimSpace(dirtyCustomer.CustomerMotherMaidenName),
		ChannelingCompany: strings.TrimSpace(dirtyCustomer.ScCompany),
		ApprovalStatus:    "9",
	}

	return customerData
}

func createLoanData(dirtyCustomer models.StagingCustomer, currentTime time.Time, custCode string) models.LoanDataTab {
	otr, err := strconv.ParseFloat(strings.TrimSpace(dirtyCustomer.LoanOtr), 32)
	if err != nil {
		log.Println(err.Error())
	}
	downPayment, err := strconv.ParseFloat(strings.TrimSpace(dirtyCustomer.LoanDownPayment), 64)
	if err != nil {
		log.Println(err.Error())
	}
	loanLoanAmountChanneling, err := strconv.ParseFloat(strings.TrimSpace(dirtyCustomer.LoanLoanAmountChanneling), 64)
	if err != nil {
		log.Println(err.Error())
	}
	interestFlat, err := strconv.ParseFloat(strings.TrimSpace(dirtyCustomer.LoanInterestFlatChanneling), 32)
	if err != nil {
		log.Println(err.Error())
	}
	interestEffective, err := strconv.ParseFloat(strings.TrimSpace(dirtyCustomer.LoanInterestEffectiveChanneling), 32)
	if err != nil {
		log.Println(err.Error())
	}
	effectivePaymentType, err := strconv.ParseInt(strings.TrimSpace(dirtyCustomer.LoanEffectivePaymentType), 10, 8)
	if err != nil {
		log.Println(err.Error())
	}
	loanMonthlyPaymentChanneling, err := strconv.ParseFloat(strings.TrimSpace(dirtyCustomer.LoanMonthlyPaymentChanneling), 64)
	if err != nil {
		log.Println(err.Error())
	}

	loanData := models.LoanDataTab{
		Custcode:             strings.TrimSpace(custCode),
		Branch:               strings.TrimSpace(dirtyCustomer.ScBranchCode),
		OTR:                  otr,
		DownPayment:          downPayment,
		LoanAmount:           loanLoanAmountChanneling,
		LoanPeriod:           strings.TrimSpace(dirtyCustomer.LoanLoanPeriodChanneling),
		InterestType:         1,
		InterestFlat:         float32(interestFlat),
		InterestEffective:    float32(interestEffective),
		EffectivePaymentType: uint8(effectivePaymentType),
		AdminFee:             30,
		MonthlyPayment:       loanMonthlyPaymentChanneling,
		InputDate:            dirtyCustomer.ScCreateDate,
		LastModified:         currentTime,
		ModifiedBy:           "system",
		Inputdate:            dirtyCustomer.ScCreateDate,
		InputBy:              "system",
		Lastmodified:         currentTime,
		Modifiedby:           "system",
	}

	return loanData
}

func createVehicleData(dirtyCustomer models.StagingCustomer, currentTime time.Time, custCode string) models.VehicleDataTab {
	vehicleType, err := strconv.ParseUint(strings.TrimSpace(dirtyCustomer.VehicleType), 10, 8)
	if err != nil {
		log.Println(err.Error())
	}
	vehicleStatus, err := strconv.ParseUint(strings.TrimSpace(dirtyCustomer.VehicleStatus), 10, 8)
	if err != nil {
		log.Println(err.Error())
	}
	vehicleDealerID, err := strconv.ParseUint(strings.TrimSpace(dirtyCustomer.VehicleDealerID), 10, 8)
	if err != nil {
		log.Println(err.Error())
	}
	vehicleTglStnk, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(dirtyCustomer.VehicleTglStnk))
	if err != nil {
		log.Println(err.Error())
	}
	vehicleTglBpkb, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(dirtyCustomer.VehicleTglBpkb))
	if err != nil {
		log.Println(err.Error())
	}
	collateralTypeID, err := strconv.ParseUint(strings.TrimSpace(dirtyCustomer.CollateralTypeID), 10, 8)
	if err != nil {
		log.Println(err.Error())
	}

	vehicleData := models.VehicleDataTab{
		Custcode:       strings.TrimSpace(custCode),
		Brand:          uint(vehicleType),
		Type:           strings.TrimSpace(dirtyCustomer.VehicleBrand),
		Year:           strings.TrimSpace(dirtyCustomer.VehicleYear),
		Golongan:       1,
		Jenis:          strings.TrimSpace(dirtyCustomer.VehicleJenis),
		Status:         uint8(vehicleStatus),
		Color:          strings.TrimSpace(dirtyCustomer.VehicleColor),
		PoliceNo:       strings.TrimSpace(dirtyCustomer.VehiclePoliceNo),
		EngineNo:       strings.TrimSpace(dirtyCustomer.VehicleEngineNo),
		ChasisNo:       strings.TrimSpace(dirtyCustomer.VehicleChasisNo),
		BPKB:           strings.TrimSpace(dirtyCustomer.VehicleBpkb),
		RegisterNo:     "1",
		STNK:           strings.TrimSpace(dirtyCustomer.VehicleStnk),
		StnkAddress1:   "",
		StnkAddress2:   "",
		StnkCity:       "",
		DealerID:       uint(vehicleDealerID),
		Inputdate:      currentTime,
		Inputby:        "system",
		Lastmodified:   currentTime,
		Modifiedby:     "system",
		TglStnk:        vehicleTglStnk,
		TglBpkb:        vehicleTglBpkb,
		TglPolis:       currentTime,
		PolisNo:        strings.TrimSpace(dirtyCustomer.VehiclePoliceNo),
		CollateralID:   collateralTypeID,
		Ketagunan:      "",
		AgunanLbu:      "",
		Dealer:         strings.TrimSpace(dirtyCustomer.VehicleDealer),
		AddressDealer1: strings.TrimSpace(dirtyCustomer.VehicleAddressDealer1),
		AddressDealer2: strings.TrimSpace(dirtyCustomer.VehicleAddressDealer2),
		CityDealer:     strings.TrimSpace(dirtyCustomer.VehicleCityDealer),
	}

	return vehicleData
}

func createStagingError(dirtyCustomer models.StagingCustomer, currentTime time.Time, errDesc string) models.StagingError {
	stagingError := models.StagingError{
		SeReff:       strings.TrimSpace(dirtyCustomer.ScReff),
		SeCreateDate: dirtyCustomer.ScCreateDate,
		BranchCode:   strings.TrimSpace(dirtyCustomer.ScBranchCode),
		Company:      strings.TrimSpace(dirtyCustomer.ScCompany),
		Ppk:          strings.TrimSpace(dirtyCustomer.CustomerPpk),
		Name:         strings.TrimSpace(dirtyCustomer.CustomerName),
		ErrorDesc:    strings.TrimSpace(errDesc),
	}

	return stagingError
}
