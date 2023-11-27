package handler

import (
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/service"

	"github.com/gin-gonic/gin"
)

type transactionHistoryHandler struct {
	transactionHistoryService service.TransactionHistoryService
}

func NewTransactionHistoryHandler(transactionHistoryService service.TransactionHistoryService) transactionHistoryHandler {
	return transactionHistoryHandler{
		transactionHistoryService: transactionHistoryService,
	}
}

func (th *transactionHistoryHandler) CreateTransaction(ctx *gin.Context) {

	addRequest := &dto.TransactionRequest{}

	if err := ctx.ShouldBindJSON(addRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	userData := ctx.MustGet("userData").(entity.User)

	response, err := th.transactionHistoryService.CreateTransaction(userData.Id, addRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.Code, response)
}

func (th *transactionHistoryHandler) GetTransactionWithProducts(ctx *gin.Context) {

	userData := ctx.MustGet("userData").(entity.User)

	response, err := th.transactionHistoryService.GetTransactionWithProducts(userData.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.Code, response)
}

func (th *transactionHistoryHandler) GetTransactionWithProductsAndUser(ctx *gin.Context) {
	response, err := th.transactionHistoryService.GetTransactionWithProductsAndUser()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.Code, response)
}
