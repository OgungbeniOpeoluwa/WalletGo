package controlers

import (
	"WalletService/dtos/request"
	"WalletService/dtos/response"
	"WalletService/services"
	"WalletService/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PaymentServiceController struct {
	walletService services.WalletService
}

func NewPaymentServiceController() *PaymentServiceController {

	return &PaymentServiceController{walletService: services.NewWalletServiceImpl()}
}

func (p *PaymentServiceController) MonifyWebhook(c *gin.Context) {
	req := request.MonnifyWebhookRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendBadRequestError(c, util.ErrMakingPostRequest)
		return
	}
	go func() {
		p.walletService.UpdateTransaction(req.EventData.PaymentReference, req.EventType)
	}()

	c.JSON(http.StatusOK, map[string]any{"status": true})

}

func (p *PaymentServiceController) PaystackWebooks(c *gin.Context) {
	req := request.PaystackWebhook{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendBadRequestError(c, err)
		return
	}
	go func() {
		p.walletService.UpdateTransaction(req.Data.Reference, req.Event)

	}()
	response.SendSuccess(c, "successful")

}
