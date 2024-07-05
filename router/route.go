package router

import (
	"WalletService/controlers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/webhook", controlers.NewPaymentServiceController().MonifyWebhook)
}
