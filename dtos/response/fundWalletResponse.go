package response

type FundWalletResponse struct {
	PaymentLink   string
	TransactionId string
}

func NewFundWalletResponse(paymentLink string, transactionId string) *FundWalletResponse {
	return &FundWalletResponse{PaymentLink: paymentLink, TransactionId: transactionId}
}
