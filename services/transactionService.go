package services

import (
	"WalletService/data/models"
	"WalletService/data/repositories"
	"WalletService/dtos/request"
	"WalletService/util"
	"github.com/google/uuid"
)

type TransactionService interface {
	CreateTransaction(req *request.CreateTransactionRequest) (*models.Transaction, error)
	GetById(id uuid.UUID) (*models.Transaction, error)
	UpdateTransaction(req *request.UpdateTransactionStatusRequest) (*models.Transaction, error)
	GetAllTransactionsByWallet(wallet uint) (*[]models.Transaction, error)
}

type TransactionServiceImpl struct {
	repository *repositories.TransactionRepositoryImpl
}

func NewTransactionServiceImpl() *TransactionServiceImpl {
	return &TransactionServiceImpl{
		repository: repositories.NewTransactionRepositoryImpl()}
}

func (receiver *TransactionServiceImpl) CreateTransaction(req *request.CreateTransactionRequest) (*models.Transaction, error) {
	respository := repositories.NewTransactionRepositoryImpl()
	transaction := mapTransaction(*req)
	createTransaction, err := respository.Save(*transaction)
	if err != nil {
		return &createTransaction, err
	}
	return &createTransaction, err

}

func (receiver *TransactionServiceImpl) GetById(id uuid.UUID) (*models.Transaction, error) {
	transaction, err := receiver.repository.FindById(id)
	if err != nil {
		err = util.ErrFetching
		return &transaction, err
	}
	return &transaction, err

}

func mapTransaction(req request.CreateTransactionRequest) *models.Transaction {
	transaction := &models.Transaction{
		Amount:        req.Amount,
		Description:   req.Description,
		PaymentStatus: req.PaymentStatus,
		PaymentType:   req.PaymentType,
		WalletID:      req.WalletId,
	}
	return transaction
}
func (receiver *TransactionServiceImpl) UpdateTransaction(req *request.UpdateTransactionStatusRequest) (*models.Transaction, error) {
	parseUuid, err := uuid.Parse(req.Id)
	if err != nil {
		err = util.ErrInvalidId
		return nil, err
	}
	transaction, err := receiver.repository.FindById(parseUuid)
	if err != nil {
		err = util.ErrFetching
		return nil, err
	}
	transaction.PaymentStatus = req.Status
	updateTransaction, err := receiver.repository.Save(transaction)
	if err != nil {
		return nil, err
	}
	return &updateTransaction, err
}

func (receiver *TransactionServiceImpl) GetAllTransactionsByWallet(walletId uint) (*[]models.Transaction, error) {
	transaction, err := receiver.repository.GetAllBy("wallet_id", walletId)
	if err != nil {
		err = util.ErrBadRequest
		return &transaction, err
	}
	return &transaction, err
}
