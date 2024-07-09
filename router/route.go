package router

import (
	"WalletService/controlers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/monify/webhook", controlers.NewPaymentServiceController().MonifyWebhook)
	router.POST("/paystack/webhook", controlers.NewPaymentServiceController().PaystackWebooks)
	walletGroup := router.Group("/api/v1/wallet")
	walletGroup.POST("/create", controlers.NewWalletController().CreateAccount)
	walletGroup.POST("/initialize-transaction", controlers.NewWalletController().InitializeTransactions)
	walletGroup.GET("/transactions/:accountNumber", controlers.NewWalletController().GetAllTransactions)
	walletGroup.GET("/balance/accountNumber", controlers.NewWalletController().GetBalance)
}
