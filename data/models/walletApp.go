package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	AccountNumber string
	Pin           string
	Balance       float64
	UserId        uint
	Password      string
}
