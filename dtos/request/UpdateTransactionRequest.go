package request

type UpdateTransactionStatusRequest struct {
	Status string `json:"status"`
	Id     string `json:"id"`
}

func NewUpdateTransactionStatusRequest(status string, id string) *UpdateTransactionStatusRequest {
	return &UpdateTransactionStatusRequest{Status: status, Id: id}
}
