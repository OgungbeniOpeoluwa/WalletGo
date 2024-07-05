package request

type FundWalletRequest struct {
	PaymentMeans  string  `json:"payment_means"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	AccountNumber string  `json:"account_number"`
	Description   string  `json:"description"`
}
