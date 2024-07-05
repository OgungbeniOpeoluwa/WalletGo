package test

import (
	"WalletService/data/models"
	"WalletService/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func MockDatabaseConnection() (*gorm.DB, error) {
	//dbs, err := openDatabase()
	//createDatabase(dbs)
	dsn := fmt.Sprintf("host=localhost user=postgres password=opemip@1 dbname=walletTest port=5432 sslmode=disable TimeZone=Asia/Shanghai")

	dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorLogger(err)
	}
	err = dbs.AutoMigrate(&models.User{}, &models.Wallet{})
	if err != nil {
		logger.ErrorLogger(err)
	}
	return dbs, err
}

func createDatabase(dbs *gorm.DB) {
	err := dbs.Exec("CREATE DATABASE walletTest").Error
	if err != nil {
		logger.ErrorLogger(err)
	}
}

func openDatabase() (*gorm.DB, error) {
	check := fmt.Sprintf("host=localhost user=postgres password=opemip@1 port=5432 sslmode=disable")
	dbs, err := gorm.Open(postgres.Open(check), &gorm.Config{})
	if err != nil {
		logger.ErrorLogger(err)
	}
	return dbs, err
}

func DropDatabaseAfterTesting() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=opemip@1 port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to postgres server")
	}
	err = db.Exec("DROP DATABASE IF EXISTS walletTest").Error
	if err != nil {
		panic(fmt.Sprintf("failed to drop database: %v", err))
	}
	log.Println("Database dropped successfully")

}
