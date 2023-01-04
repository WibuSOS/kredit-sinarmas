package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sinarmas/kredit-sinarmas/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	migrate_flag := os.Getenv("AUTO_MIGRATE")
	seed_flag := os.Getenv("SEED_DB")

	db, env, err := callDb()

	if err != nil {
		return nil, err
	}

	db, err = checkDbConn(db)
	if err != nil {
		return nil, err
	}

	if migrate_flag == "true" {
		db, err = migrateDb(db)
	}
	if err != nil {
		return nil, errorDbConn(err)
	}

	if seed_flag == "true" && env != "PROD" {
		seedDb(db)
	}

	return db, nil
}

func errorDbConn(err error) error {
	return fmt.Errorf("failed to connect database: %w", err)
}

func callDb() (*gorm.DB, string, error) {
	var db *gorm.DB
	var err error
	env := os.Getenv("ENVIRONMENT")

	// if env == "PROD" {
	// 	dbUrl := os.Getenv("DATABASE_URL")
	// 	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	// }

	// if env == "STAGING" {
	// 	dbUrl := os.Getenv("DATABASE_URL")
	// 	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	// 	db.Exec("DROP TABLE _______")
	// }

	// if env == "TEST" {
	// 	db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	// 	if err != nil {
	// 		return nil, env, errorDbConn(err)
	// 	}

	// 	err = db.Exec("PRAGMA foreign_keys = ON", nil).Error
	// }

	// if err != nil {
	// 	return nil, env, errorDbConn(err)
	// }

	// if db != nil {
	// 	log.Println("Call DB success")
	// 	return db, env, nil
	// }

	db, err = callDbDev()

	if err != nil {
		return nil, env, errorDbConn(err)
	}

	log.Println("Call DB success")
	return db, env, nil
}

func callDbDev() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Open DB Root only for creating the intended DB
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_SUPER_USER"), os.Getenv("DB_SUPER_PASSWORD"), os.Getenv("DB_ROOT"))
	dbRoot, errRoot := gorm.Open(postgres.Open(config), &gorm.Config{})

	if errRoot != nil {
		return nil, errorDbConn(errRoot)
	}

	// Implicitly silences error in case the intended DB already exists
	dbRoot.Exec(fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("DB_NAME")))

	// Close DB Root
	sqlDbRoot, errRoot := dbRoot.DB()
	if errRoot != nil {
		return nil, errRoot
	}
	errRoot = sqlDbRoot.Close()
	if errRoot != nil {
		return nil, errRoot
	}

	// Open the intended DB
	config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, errorDbConn(err)
	}

	db.Exec("DROP TABLE branch_tab")
	db.Exec("DROP TABLE mst_company_tab")
	db.Exec("DROP TABLE customer_data_tab")
	db.Exec("DROP TABLE loan_data_tab")

	log.Println("Call DB Dev success")

	return db, nil
}

func checkDbConn(db *gorm.DB) (*gorm.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errorDbConn(err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, errorDbConn(err)
	}

	log.Println("Check DB connection success")
	return db, nil
}

func migrateDb(db *gorm.DB) (*gorm.DB, error) {
	if err := db.AutoMigrate(models.BranchTab{}, models.MstCompanyTab{}, models.CustomerDataTab{}, models.LoanDataTab{}); err != nil {
		return nil, errorDbConn(err)
	}

	log.Println("Migrate DB success")
	return db, nil
}

func seedDb(db *gorm.DB) {
	// pb, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 8)
	// newUsers := []models.Users{
	// 	{Nama: "Penjual", Role: "consumer", NoHp: "+6285775066878", Email: "penjual@custom.com", Password: string(pb), NoRek: "1234567890"},
	// 	{Nama: "Pembeli", Role: "consumer", NoHp: "+6281586173213", Email: "pembeli@custom.com", Password: string(pb), NoRek: "0987654321"},
	// }
	// seedTable(db, &models.Users{}, &newUsers)
}

func seedTable(db *gorm.DB, table any, newRecords any) {
	if !db.Migrator().HasTable(table) {
		return
	}

	if err := db.First(table).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		res := db.Create(newRecords)
		if res.Error != nil {
			log.Println(res.Error.Error())
		}
	}
}
