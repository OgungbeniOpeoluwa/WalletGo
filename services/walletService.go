package services

import (
	"WalletService/data/models"
	"WalletService/data/repositories"
	"WalletService/dtos/request"
	"WalletService/dtos/response"
	"WalletService/services/paymentService"
	"WalletService/util"
	"fmt"
)

type WalletService interface {
	CreateAccount(req request.CreateAccountRequest) (string, error)
	InitializeTransaction(req request.FundWalletRequest) (*response.FundWalletResponse, error)
	GetAllTransactions(accountNumber string) ([]*response.TransactionResponse, error)
	GetBalance(accountNumber string) (*float64, error)
}

type WalletServiceImpl struct {
	repository         *repositories.WalletRepositoryImpl
	userService        UserService
	transactionService TransactionService
}

func NewWalletServiceImpl() *WalletServiceImpl {
	return &WalletServiceImpl{
		repository:         repositories.NewWalletRepositoryImpl(),
		userService:        NewUserServiceImpl(),
		transactionService: NewTransactionServiceImpl(),
	}
}

func (receiver *WalletServiceImpl) CreateAccount(req request.CreateAccountRequest) (string, error) {
	user := request.CreateUserRequest{
		LastName:    req.LastName,
		FirstName:   req.FirstName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}
	users, err := receiver.userService.CreateUser(user)
	if err != nil {
		return "", err
	}

	account := models.Wallet{
		AccountNumber: user.PhoneNumber,
		UserId:        users.ID,
		Password:      req.Password,
	}
	save, err := receiver.repository.Save(account)
	if err != nil {
		return "", err
	}
	return save.AccountNumber, err

}

func (receiver *WalletServiceImpl) InitializeTransaction(req request.FundWalletRequest) (*response.FundWalletResponse, error) {
	var user *models.User
	user, err := receiver.userService.GetUserByPhoneNumber(req.AccountNumber)
	wallet, err := receiver.repository.GetBy("user_id", user.ID)
	if err != nil {
		err = util.ErrFetching
		return nil, err
	}
	transaction, err2 := mapCreateTransactionName(req, wallet, receiver)

	if err2 != nil {
		return nil, err2
	}
	paymentRequest := *request.NewPaymentServiceRequest(req.Amount, user.Email,
		req.Currency, req.PaymentMeans, req.Description, user.FirstName, transaction.ID.String())

	paymentTransaction, err := paymentService.InitiateTransaction(paymentRequest)

	if err != nil {
		err, walletResponse, err3 := updateFailedTransaction(receiver, transaction)
		if err3 != nil {
			return walletResponse, err3
		}
		return nil, err
	}
	return paymentTransaction, err
}

func (receiver *WalletServiceImpl) GetAllTransactions(s string) ([]*response.TransactionResponse, error) {
	wallet, err := receiver.repository.GetBy("account_number", s)
	if err != nil {
		err = fmt.Errorf("wallet %v", util.ErrFetching)
		return nil, err
	}
	transactions, err := receiver.transactionService.GetAllTransactionsByWallet(wallet.ID)
	if err != nil {
		return nil, err
	}
	transactionResponse := mapTransactionResponse(transactions)
	return transactionResponse, nil
}

func (receiver *WalletServiceImpl) GetBalance(s string) (*float64, error) {
	account, err := receiver.repository.GetBy("account_number", s)
	if err != nil {
		err = fmt.Errorf("wallet %v", util.ErrFetching)
		return nil, err
	}
	return &account.Balance, err

}

func mapTransactionResponse(transactions *[]models.Transaction) []*response.TransactionResponse {
	var transactionResponse []*response.TransactionResponse
	for _, transaction := range *transactions {
		maps := response.NewTransactionResponse(
			transaction.ID,
			transaction.Amount,
			transaction.Description,
			transaction.PaymentStatus,
			transaction.PaymentType)
		transactionResponse = append(transactionResponse, maps)
	}
	return transactionResponse
}

func mapCreateTransactionName(req request.FundWalletRequest, wallet models.Wallet, receiver *WalletServiceImpl) (*models.Transaction, error) {
	createTransaction := &request.CreateTransactionRequest{
		Amount:        req.Amount,
		Description:   req.Description,
		WalletId:      wallet.ID,
		PaymentType:   util.Credit,
		PaymentStatus: util.Processing,
	}
	transaction, err := receiver.transactionService.CreateTransaction(createTransaction)
	if err != nil {
		err = util.ErrProcessingPayment
		return nil, err
	}
	return transaction, nil
}

func updateFailedTransaction(receiver *WalletServiceImpl, transaction *models.Transaction) (error, *response.FundWalletResponse, error) {
	_, err := receiver.transactionService.UpdateTransaction(request.NewUpdateTransactionStatusRequest(
		util.Failed, transaction.ID))
	if err != nil {
		return nil, nil, err
	}
	return err, nil, nil
}
