package request

type CreateUserRequest struct {
	FirstName   string `json:"firstname"`
	Email       string `json:"email"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}
