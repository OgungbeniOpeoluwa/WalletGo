package controlers

import (
	"WalletService/dtos/request"
	"WalletService/dtos/response"
	"WalletService/services"
	"github.com/gin-gonic/gin"
)

type WalletController struct {
	walletService services.WalletService
}

func NewWalletController() *WalletController {
	return &WalletController{walletService: services.NewWalletServiceImpl()}
}

func (C *WalletController) CreateAccount(c *gin.Context) {
	req := request.CreateAccountRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendBadRequestError(c, err.Error())
		return
	}
	account, err := C.walletService.CreateAccount(req)
	if err != nil {
		response.SendUnprocessableEntity(c, err.Error())
		return
	}
	response.SendSuccess(c, account)

}

func (C *WalletController) InitializeTransactions(c *gin.Context) {
	req := request.FundWalletRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendBadRequestError(c, err.Error())
		return
	}
	account, err := C.walletService.InitializeTransaction(req)
	if err != nil {
		response.SendUnprocessableEntity(c, err.Error())
		return
	}
	response.SendSuccess(c, account)
}

func (C *WalletController) GetAllTransactions(c *gin.Context) {
	req := c.Param("accountNumber")
	transactions, err := C.walletService.GetAllTransactions(req)
	if err != nil {
		response.SendUnprocessableEntity(c, err)
		return
	}
	response.SendSuccess(c, transactions)
}

func (C *WalletController) GetBalance(c *gin.Context) {
	req := c.Param("accountNumber")
	balance, err := C.walletService.GetBalance(req)
	if err != nil {
		response.SendUnprocessableEntity(c, err.Error())
		return
	}
	response.SendSuccess(c, *balance)
}
