package request

type CreatePaymentServiceRequest struct {
	Amount       float64 `json:"amount"`
	Email        string  `json:"email"`
	Currency     string  `json:"currency"`
	PaymentMeans string  `json:"payment_means"`
	Description  string  `json:"description"`
	CustomerName string  `json:"customer_name"`
	Reference    string  `json:"reference"`
}

func NewPaymentServiceRequest(amount float64, email string, currency string, paymentMeans string, description string, customerName string,
	reference string) *CreatePaymentServiceRequest {
	return &CreatePaymentServiceRequest{Amount: amount, Email: email, Currency: currency,
		PaymentMeans: paymentMeans, Description: description, CustomerName: customerName,
		Reference: reference}
}
