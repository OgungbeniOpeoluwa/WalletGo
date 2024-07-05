package util

import "errors"

var (
	ErrMakingPostRequest = errors.New("error making post httpRequest")
	ErrDecodingResponse  = errors.New("error Decoding response")
	ErrPaymentMedium     = errors.New("payment medium doesnt exist")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrProcessingPayment = errors.New("err processing transaction")
	ErrCreatingUser      = errors.New("failed to create user")
	ErrFetching          = errors.New("doesn't exist")
	ErrAlreadyExit       = errors.New("account already exist")
	ErrBadRequest        = errors.New("bad request")
)
