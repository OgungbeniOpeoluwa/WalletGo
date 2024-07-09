package request

type Data struct {
	Reference string  `json:"reference"`
	Amount    float64 `json:"amount"`
}

type PaystackWebhook struct {
	Data  Data   `json:"data"`
	Event string `json:"event"`
}
