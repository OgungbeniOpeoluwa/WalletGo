package repositories

import "WalletService/data/models"

type UserRepositoryImpl struct {
	BaseRepositoryInterface[models.User]
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	baseRepo := NewBaseRepository[models.User]()
	return &UserRepositoryImpl{baseRepo}
}
