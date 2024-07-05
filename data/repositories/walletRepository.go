package repositories

import "WalletService/data/models"

type WalletRepositoryImpl struct {
	BaseRepositoryInterface[models.Wallet]
}

func NewWalletRepositoryImpl() *WalletRepositoryImpl {
	baseRepo := NewBaseRepository[models.Wallet]()
	return &WalletRepositoryImpl{baseRepo}
}
