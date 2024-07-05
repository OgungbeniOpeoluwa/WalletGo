package response

import "github.com/google/uuid"

type TransactionResponse struct {
	Id            uuid.UUID `json:"id"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	PaymentStatus string    `json:"paymentStatus"`
	PaymentType   string    `json:"paymentType"`
}

func NewTransactionResponse(id uuid.UUID, amount float64, description string, paymentStatus string, paymentType string) *TransactionResponse {
	return &TransactionResponse{Id: id, Amount: amount, Description: description, PaymentStatus: paymentStatus, PaymentType: paymentType}
}
