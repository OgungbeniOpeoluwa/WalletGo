package response

type MonifyResponse struct {
	RequestSuccessful bool         `json:"requestSuccessful"`
	ResponseMessage   string       `json:"responseMessage"`
	ResponseCode      string       `json:"responseCode"`
	ResponseBody      ResponseBody `json:"responseBody"`
}

type ResponseBody struct {
	TransactionReference string `json:"transactionReference"`
	PaymentReference     string `json:"paymentReference"`
	CheckOutUrl          string `json:"checkoutUrl"`
}
