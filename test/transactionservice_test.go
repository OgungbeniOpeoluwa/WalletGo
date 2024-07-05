package test

import (
	"WalletService/dtos/request"
	"WalletService/services"
	"WalletService/util"
	"WalletService/util/config"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	config.Load("../.env ")
	transaction := services.NewTransactionServiceImpl()
	createTransaction := request.CreateTransactionRequest{
		Amount:        1000,
		PaymentType:   util.Credit,
		PaymentStatus: util.Pending,
	}
	returns, err := transaction.CreateTransaction(&createTransaction)
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
	}
	assert.Equal(t, returns.Amount, float64(1000))

}

func TestFindTransactionByIdIfDoesntExistThrowDoesntExist(t *testing.T) {
	config.Load("../.env ")
	transaction := services.NewTransactionServiceImpl()
	_, err := transaction.GetById(uuid.New())
	err2 := util.ErrFetching
	if err == nil {
		t.Errorf("expected transaction doesn't exist")
	}
	assert.Equal(t, err, err2)

}
func TestFindTransactionByIdAndTransactionIsReturn(t *testing.T) {
	config.Load("../.env ")
	transaction := services.NewTransactionServiceImpl()
	createTransaction := request.CreateTransactionRequest{
		Amount:        1000,
		Description:   "My hair",
		PaymentType:   util.Credit,
		PaymentStatus: util.Pending,
	}
	returns, err := transaction.CreateTransaction(&createTransaction)
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
	}
	transactions, err := transaction.GetById(returns.ID)
	if err != nil {
		t.Errorf("expected nil but got %v:", err)
	}
	assert.Equal(t, transactions.ID, returns.ID)

}

func TestUpdateTransactionStatusFromPendingToFailed(t *testing.T) {
	config.Load("../.env ")
	transaction := services.NewTransactionServiceImpl()
	createTransaction := request.CreateTransactionRequest{
		Amount:        1000,
		Description:   "My hair",
		PaymentType:   util.Credit,
		PaymentStatus: util.Pending,
	}
	returns, err := transaction.CreateTransaction(&createTransaction)
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
		return
	}

	updateTransaction, err := transaction.UpdateTransaction(&request.UpdateTransactionStatusRequest{
		Id:     returns.ID,
		Status: util.Failed,
	})
	if err != nil {
		t.Errorf("expected nil but got %v:", err)
		return
	}
	assert.Equal(t, updateTransaction.ID, returns.ID)
	assert.NotEqual(t, updateTransaction.PaymentStatus, returns.PaymentStatus)

}
