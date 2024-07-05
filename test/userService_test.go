package test

import (
	"WalletService/dtos/request"
	"WalletService/services"
	"WalletService/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	createRequest := request.CreateUserRequest{
		FirstName:   "opeoluwa",
		LastName:    "tola",
		PhoneNumber: "07066221008",
		Email:       "ope@gmail.com",
	}
	users := services.NewUserServiceImpl()
	user, err := users.CreateUser(createRequest)
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
	}
	assert.Equal(t, createRequest.FirstName, user.FirstName)
}

func TestThatUserThatWithTheSamePhoneNumberCantRegister(t *testing.T) {
	createRequest := request.CreateUserRequest{
		FirstName:   "opeoluwa",
		LastName:    "tola",
		PhoneNumber: "07066221008",
		Email:       "ope@gmail.com",
	}
	users := services.NewUserServiceImpl()
	user, err := users.CreateUser(createRequest)
	err2 := util.ErrAlreadyExit
	if err == nil {
		t.Errorf("expected %v but got %v", user, err)
	}
	assert.Equal(t, err2, err)
}

func TestGetUserByEmail(t *testing.T) {
	createRequest := request.CreateUserRequest{
		FirstName:   "opeoluwa",
		LastName:    "tola",
		PhoneNumber: "07066827808",
		Email:       "ope@gmail.com",
	}
	users := services.NewUserServiceImpl()
	user, err := users.CreateUser(createRequest)
	if err != nil {
		t.Errorf("expected %v but got %v", user, err)
		return
	}
	getUser, err := users.GetUserByEmail("ope@gmail.com")
	if err != nil {
		t.Errorf("expected %v but got %v", nil, err)
	}
	assert.Equal(t, getUser.Email, "ope@gmail.com")
}
