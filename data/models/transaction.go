package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            uuid.UUID `gorm:"primaryKey"`
	Description   string
	Amount        float64
	RecipientName string
	WalletID      uint
	PaymentType   string
	PaymentStatus string
}

func (receiver *Transaction) BeforeCreate(db *gorm.DB) (err error) {
	receiver.ID = uuid.New()
	return
}
