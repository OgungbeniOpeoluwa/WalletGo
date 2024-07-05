package db

import (
	"WalletService/data/models"
	"WalletService/logger"
	"WalletService/util/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

func DbConnection() (dbs *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DatabaseHost, config.DatabaseUsername, config.DatabasePassword, config.DatabaseName, config.DatabasePort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorLogger(err)

	}
	err = db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})
	if err != nil {
		logger.ErrorLogger(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	return db, err

}
