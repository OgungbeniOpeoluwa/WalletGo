package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint   `json:"primaryKey"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
