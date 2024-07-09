package request

type MonnifyWebhookRequest struct {
	EventType string    `json:"eventType"`
	EventData EventData `json:"eventData"`
}

type EventData struct {
	PaymentReference string `json:"paymentReference"`
}
