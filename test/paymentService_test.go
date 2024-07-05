package test

import (
	"WalletService/dtos/request"
	"WalletService/services/paymentService"
	"WalletService/util/config"
	"fmt"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

func TestFundWalletUsingPaystack(t *testing.T) {
	config.Load("../.env")
	paymentRequest := *request.NewPaymentServiceRequest(1200, "weyotok394@dcbin.com", "NGN", "Paystack",
		"payment for school", "shola", "76763gebyyhy3")
	result, err := paymentService.InitiateTransaction(paymentRequest)
	if err != nil {
		t.Errorf("expected %v but got %v", result, err)
	}
	log.Println(result)
}

func TestInitializeTransactionThrowAnErrorIfPaymentMeansDoesNotExist(t *testing.T) {
	config.Load("../.env")
	payment := *request.NewPaymentServiceRequest(1200, "weyotok394@dcbin.com", "NGN", "Kuda",
		"payment for drink", "shola", "8847738736776674")
	_, err := paymentService.InitiateTransaction(payment)
	err2 := fmt.Errorf("payment medium doesnt exist")
	if err == nil {
		t.Errorf("expected %v but got %v", err2, err)
	}
	assert.Equal(t, err2, err)

}

func TestThatIfAmountIsLesserThan10ErrorIsThrown(t *testing.T) {
	config.Load("../.env")
	payment := *request.NewPaymentServiceRequest(5, "weyotok394@dcbin.com", "NGN", "paystack",
		"payment for hair", "shola", "223411455-joko")
	_, err := paymentService.InitiateTransaction(payment)
	err2 := fmt.Errorf("invalid amount")
	if err == nil {
		t.Errorf("expected %v but got %v", err2, err)
	}
	assert.Equal(t, err2, err)
}

func TestThatMoneyCanBeSendThroughAnotherPaymentMedium(t *testing.T) {
	config.Load("../.env")
	payment := *request.NewPaymentServiceRequest(1200, "weyotok394@dcbin.com", "NGN", "monnify",
		"payment for drink", "shayo", "23224rdd-oooo")
	result, err := paymentService.InitiateTransaction(payment)
	err2 := fmt.Errorf("nil")
	if err != nil {
		t.Errorf("expected %v but got %v", err2, err)
	}
	log.Println(result)
}
