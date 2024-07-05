package request

type MonnifyWebhookRequest struct {
	EventType string    `json:"eventType"`
	EventData EventData `json:"eventData"`
}

type EventData struct {
	AmountPaid         float64  `json:"amountPaid"`
	PaymentDescription string   `json:"paymentDescription"`
	PaymentReference   string   `json:"paymentReference"`
	Customer           Customer `json:"customer"`
}

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
