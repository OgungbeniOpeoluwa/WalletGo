package request

import "github.com/google/uuid"

type UpdateTransactionStatusRequest struct {
	Status string    `json:"status"`
	Id     uuid.UUID `json:"id"`
}

func NewUpdateTransactionStatusRequest(status string, id uuid.UUID) *UpdateTransactionStatusRequest {
	return &UpdateTransactionStatusRequest{Status: status, Id: id}
}
