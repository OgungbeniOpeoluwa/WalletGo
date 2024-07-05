package util

import (
	"WalletService/dtos/request"
	"WalletService/dtos/response"
	"WalletService/util/config"
	"strings"
)

func CreatePaystackPaymentRequest(serviceRequest request.CreatePaymentServiceRequest) map[string]any {
	var amount = serviceRequest.Amount
	if strings.ToUpper(serviceRequest.Currency) == Naira {
		amount *= 100
	}
	return map[string]any{
		"amount":    amount,
		"email":     serviceRequest.Email,
		"currency":  serviceRequest.Currency,
		"reference": serviceRequest.Reference,
	}

}
func CreateMonifyPaymentRequest(requestBody request.CreatePaymentServiceRequest) map[string]any {
	return map[string]any{
		"amount":             requestBody.Amount,
		"customerEmail":      requestBody.Email,
		"currencyCode":       requestBody.Currency,
		"customerName":       requestBody.CustomerName,
		"paymentDescription": requestBody.Description,
		"paymentReference":   requestBody.Reference,
		"contractCode":       config.ContractCode,
	}

}
func IntializeTransactionResponse(url string, transactionID string) *response.FundWalletResponse {
	return response.NewFundWalletResponse(url, transactionID)
}
