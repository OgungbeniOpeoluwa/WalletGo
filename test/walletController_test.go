package test

import (
	"WalletService/controlers"
	"WalletService/dtos/request"
	"WalletService/util/config"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	config.Load("../.env")
	var accountHandler = controlers.NewWalletController()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	t.Run("success", func(t *testing.T) {
		router.POST("/api/v1/wallet/account", accountHandler.CreateAccount)
		writer := httptest.NewRecorder()
		createAccountRequest := request.NewCreateAccountRequest("ope", "shayo", "ope1@gmail.com", "070662921008", "ayo10")
		data, _ := json.Marshal(createAccountRequest)
		requests, _ := http.NewRequest(http.MethodPost, "/api/v1/wallet/account", bytes.NewReader(data))
		requests.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(writer, requests)
		log.Println(writer.Body.String())
		assert.Equal(t, http.StatusOK, writer.Code)
	})
}

func TestInitializeTransactionToRegisterAccount(t *testing.T) {
	config.Load("../.env")
	var accountHandler = controlers.NewWalletController()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	t.Run("success", func(t *testing.T) {
		router.POST("/api/v1/wallet/initialize-transaction", accountHandler.InitializeTransactions)
		writer := httptest.NewRecorder()
		initializeRequest := request.FundWalletRequest{
			Amount:        1000,
			PaymentMeans:  "monnify",
			Description:   "Hair",
			AccountNumber: "070662921008",
			Currency:      "NGN",
		}
		data, _ := json.Marshal(initializeRequest)
		requests, _ := http.NewRequest(http.MethodPost, "/api/v1/wallet/initialize-transaction", bytes.NewReader(data))
		requests.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(writer, requests)
		log.Println(writer.Body.String())
		assert.Equal(t, http.StatusOK, writer.Code)
	})

}

func TestGetAllTansactionWithAccountNumber(t *testing.T) {
	config.Load("../.env")
	var account = controlers.NewWalletController()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	t.Run("success", func(t *testing.T) {
		router.GET("/api/v1/wallet/transactions/:accountNumber", account.GetAllTransactions)
		writer := httptest.NewRecorder()
		requests, _ := http.NewRequest(http.MethodGet, "/api/v1/wallet/transactions/070662921008", nil)
		router.ServeHTTP(writer, requests)
		log.Println(writer.Body.String())
		assert.Equal(t, http.StatusOK, writer.Code)

	})

}

func TestGetBalanceWithAccountNumber(t *testing.T) {
	config.Load("../.env")
	var account = controlers.NewWalletController()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	t.Run("success", func(t *testing.T) {
		router.GET("/api/v1/wallet/balance/:accountNumber", account.GetBalance)
		writer := httptest.NewRecorder()
		data, _ := http.NewRequest(http.MethodGet, "/api/v1/wallet/balance/070662921008", nil)
		data.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(writer, data)
		log.Println(writer.Body.String())
		assert.Equal(t, http.StatusOK, writer.Code)
	})
}
