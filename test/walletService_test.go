package test

import (
	"WalletService/dtos/request"
	"WalletService/services"
	"WalletService/util"
	"WalletService/util/config"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	createAccount := request.CreateAccountRequest{
		FirstName:   "ope",
		LastName:    "shola",
		Email:       "ope@gmail.com",
		PhoneNumber: "07066221008",
		Password:    "ope@1",
	}
	account := services.NewWalletServiceImpl()
	accountNumber, err := account.CreateAccount(createAccount)
	if err != nil {
		t.Errorf("expected %v but got %v", err, accountNumber)
	}

}

func TestThatAmountLesserThan10CantBeTransfer(t *testing.T) {
	createAccount := request.CreateAccountRequest{
		FirstName:   "ope",
		LastName:    "shola",
		Email:       "ope@gmail.com",
		PhoneNumber: "07066221098",
		Password:    "ope@1",
	}
	account := services.NewWalletServiceImpl()
	_, err := account.CreateAccount(createAccount)
	_, err = account.InitializeTransaction(request.FundWalletRequest{
		AccountNumber: "07066221098",
		Amount:        5,
		PaymentMeans:  "paystack",
		Description:   "my hair",
	})
	err2 := util.ErrInvalidAmount
	assert.Equal(t, err, err2)

}

func TestInitializeTransaction(t *testing.T) {
	config.Load("../.env")
	createAccount := request.CreateAccountRequest{
		FirstName:   "ope",
		LastName:    "shola",
		Email:       "ope3@gmail.com",
		PhoneNumber: "07063620998",
	}
	account := services.NewWalletServiceImpl()
	_, err := account.CreateAccount(createAccount)
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
		return
	}
	fund, err := account.InitializeTransaction(request.FundWalletRequest{
		AccountNumber: "07063620998",
		Amount:        1000,
		PaymentMeans:  "paystack",
		Description:   "my hair",
		Currency:      "NGN",
	})
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
	}
	log.Println(fund)

}

func TestAfterTransactionIsCreated(t *testing.T) {
	config.Load("../.env")
	createAccount := request.CreateAccountRequest{
		FirstName:   "ope",
		LastName:    "shola",
		Email:       "ope3@gmail.com",
		PhoneNumber: "07063620998",
		Password:    "ope@1",
	}
	account := services.NewWalletServiceImpl()
	_, err := account.CreateAccount(createAccount)
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
		return
	}
	result, err := account.InitializeTransaction(request.FundWalletRequest{
		AccountNumber: "07063620998",
		Amount:        1000,
		PaymentMeans:  "paystack",
		Description:   "my hair",
		Currency:      "NGN",
	})
	log.Println(result)
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
	}
	transaction, err := account.GetAllTransactions("07063620998")
	if err != nil {
		t.Errorf("expected nil but got %v", err)
		return
	}
	assert.Equal(t, len(transaction), 1)

}

func TestGetAccountBalance(t *testing.T) {
	config.Load("../.env")
	account := services.NewWalletServiceImpl()
	accountBalance, err := account.GetBalance("07063620998")
	if err != nil {
		t.Errorf("expected nil but got %v", err)
		return
	}
	assert.Equal(t, float64(0), *accountBalance)

}
