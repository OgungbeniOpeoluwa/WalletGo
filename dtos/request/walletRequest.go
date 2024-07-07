package request

type CreateAccountRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func NewCreateAccountRequest(firstName, lastName, email, phoneNumber, password string) *CreateAccountRequest {
	return &CreateAccountRequest{FirstName: firstName, LastName: lastName, Email: email, PhoneNumber: phoneNumber, Password: password}
}

type CreateAccountResponse struct {
	AccountNumber string `json:"accountNumber"`
}
