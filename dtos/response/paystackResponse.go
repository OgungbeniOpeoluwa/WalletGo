package response

type Paystack struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

type data struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}
