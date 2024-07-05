package PaystackService

import (
	"WalletService/dtos/response"
	"WalletService/logger"
	"WalletService/util"
	"WalletService/util/config"
	"WalletService/util/httpRequest"
)

type PaystackServiceImpl struct {
}

func NewImpl() *PaystackServiceImpl {
	return &PaystackServiceImpl{}
}

func (receiver *PaystackServiceImpl) FundWallet(req map[string]any) (*response.FundWalletResponse, error) {
	var r = response.Paystack{}
	key := "Bearer " + config.PaystackSecretKey
	responseBody, err := httpRequest.MakePostRequest(req, config.PaystackInitializeTransactionUrl, key)
	if err != nil {
		logger.ErrorLogger(err)
		return nil, util.ErrMakingPostRequest
	}
	r, err = httpRequest.DecodeRequestBody[response.Paystack](responseBody, r)
	if err != nil {
		logger.ErrorLogger(err)
		return nil, util.ErrDecodingResponse
	}
	requestResponse := util.IntializeTransactionResponse(r.Data.AuthorizationUrl, r.Data.Reference)
	return requestResponse, nil
}
