package controlers

import (
	"WalletService/dtos/request"
	"WalletService/dtos/response"
	"WalletService/util"
	"WalletService/util/httpRequest"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PaymentServiceController struct {
	//c *gin.Context
}

func NewPaymentServiceController() *PaymentServiceController {

	return &PaymentServiceController{}
}

func (p *PaymentServiceController) MonifyWebhook(c *gin.Context) {
	log.Println("am here: ", c.Request)
	req := request.MonnifyWebhookRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("err: ", err)
		response.SendBadRequestError(c, util.ErrMakingPostRequest)
		return
	}
	log.Println(req)
	res, err := httpRequest.DecodeRequestBody[request.MonnifyWebhookRequest](c.Request.Body, request.MonnifyWebhookRequest{})
	//if res.EventData.PaymentStatus == "SUCCESSFUL_TRANSACTION" {
	//
	//}
	if err != nil {
		response.SendBadRequestError(c, util.ErrMakingPostRequest)
		return
	}
	log.Println(res)
	c.JSON(http.StatusOK, map[string]any{"status": true})

}
