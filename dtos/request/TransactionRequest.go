package request

type CreateTransactionRequest struct {
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
	WalletId      uint    `json:"walletId"`
	PaymentType   string  `json:"paymentType"`
	PaymentStatus string  `json:"paymentStatus"`
}
