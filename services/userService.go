package services

import (
	"WalletService/data/models"
	"WalletService/data/repositories"
	"WalletService/dtos/request"
	"WalletService/util"
)

type UserService interface {
	CreateUser(req request.CreateUserRequest) (*models.User, error)
	GetUserByPhoneNumber(phonenumber string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserServiceImpl struct {
	repository *repositories.UserRepositoryImpl
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{repository: repositories.NewUserRepositoryImpl()}
}

func (i *UserServiceImpl) CreateUser(req request.CreateUserRequest) (*models.User, error) {
	_, err := i.GetUserByPhoneNumber(req.PhoneNumber)
	if err == nil {
		err = util.ErrAlreadyExit
		return nil, err
	}
	user := createUser(req)
	save, err := i.repository.Save(user)

	if err != nil {
		err = util.ErrCreatingUser
		return &save, err
	}
	return &save, err

}

func createUser(req request.CreateUserRequest) models.User {
	user := models.User{
		LastName:    req.LastName,
		FirstName:   req.FirstName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	return user

}

func (i *UserServiceImpl) GetUserByPhoneNumber(number string) (*models.User, error) {
	user, err := i.repository.GetBy("phone_number", number)
	if err != nil {
		err = util.ErrFetching
		return &user, err
	}
	return &user, err

}

func (i *UserServiceImpl) GetUserByEmail(req string) (*models.User, error) {
	user, err := i.repository.GetBy("email", req)
	if err != nil {
		err = util.ErrFetching
		return &user, err
	}
	return &user, err

}
