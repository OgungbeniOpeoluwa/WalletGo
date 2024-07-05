package paymentService

import (
	"WalletService/dtos/request"
	"WalletService/dtos/response"
	"WalletService/services/PaystackService"
	"WalletService/services/monnify"
	"WalletService/util"
	"strings"
)

type PaymentService interface {
	FundWallet(req map[string]any) (*response.FundWalletResponse, error)
}

func InitiateTransaction(request request.CreatePaymentServiceRequest) (*response.FundWalletResponse, error) {
	var dataResponse *response.FundWalletResponse
	var err error
	if request.Amount < 10 {
		err = util.ErrInvalidAmount
		return dataResponse, err
	}
	switch strings.ToUpper(request.PaymentMeans) {
	case util.Paystack:
		req := util.CreatePaystackPaymentRequest(request)
		dataResponse, err = PaystackService.NewImpl().FundWallet(req)
		if err != nil {
			return nil, err
		}
		return dataResponse, err
	case util.Monnify:
		req := util.CreateMonifyPaymentRequest(request)
		dataResponse, err = monnify.NewMonnifyServiceImpl().FundWallet(req)
		if err != nil {
			return dataResponse, err
		}
		return dataResponse, err

	default:
		err = util.ErrPaymentMedium
		return dataResponse, err
	}

}
