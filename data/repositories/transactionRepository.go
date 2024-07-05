package repositories

import (
	"WalletService/data/models"
)

type TransactionRepositoryImpl struct {
	BaseRepositoryInterface[models.Transaction]
}

func NewTransactionRepositoryImpl() *TransactionRepositoryImpl {
	baseRepo := NewBaseRepository[models.Transaction]()
	return &TransactionRepositoryImpl{baseRepo}
}

func (receiver *TransactionRepositoryImpl) FindAllTransactionByWalletId(id uint) (*[]models.Transaction, error) {
	transactions, err := receiver.GetAllBy("id", id)
	if err != nil {
		return &transactions, err
	}
	return &transactions, err
}
