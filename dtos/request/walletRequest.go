package request

type CreateAccountRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
type CreateAccountResponse struct {
	accountNumber string `json:"accountNumber"`
}
