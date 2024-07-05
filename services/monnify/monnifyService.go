package monnify

import (
	"WalletService/dtos/response"
	"WalletService/logger"
	"WalletService/util"
	"WalletService/util/config"
	"WalletService/util/httpRequest"
	"fmt"
	"log"
)

type MonnifyService struct {
}

func NewMonnifyServiceImpl() *MonnifyService {
	return &MonnifyService{}
}

func (reciever *MonnifyService) FundWallet(req map[string]any) (*response.FundWalletResponse, error) {
	var err error
	var decodeBody = response.MonifyResponse{}
	decodeKey := util.Encode(fmt.Sprintf("%s:%s", config.MonifyApiKey, config.MonifySecretKey))
	key := "Basic " + decodeKey
	jsonResponse, err := httpRequest.MakePostRequest(req, config.MonifyInitializeTransactionUrl, key)
	if err != nil {
		logger.ErrorLogger(err)
		return nil, err
	}
	decodeBody, err = httpRequest.DecodeRequestBody[response.MonifyResponse](jsonResponse, decodeBody)
	if err != nil {
		return nil, err
	}
	if decodeBody.RequestSuccessful != true {
		log.Println(decodeBody.ResponseMessage)
		return nil, util.ErrProcessingPayment
	}
	requestResponse := util.IntializeTransactionResponse(decodeBody.ResponseBody.CheckOutUrl, decodeBody.ResponseBody.TransactionReference)
	return requestResponse, nil
}
