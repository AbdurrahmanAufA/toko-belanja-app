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

// CreateNewTransaction godoc
// @Tags movies
// @Description Create New Movie Data
// @ID create-new-movie
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewMovieRequest true "request body json"
// @Success 201 {object} dto.NewMovieRequest
// @Router /movies [post]
// func (m transactionHistoryHandler) CreateTransactionHistory(c *gin.Context) {
// 	var transactionHistoryRequest dto.NewTransactionHistoryRequest

// 	if err := c.ShouldBindJSON(&transactionHistoryRequest); err != nil {
// 		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

// 		c.JSON(errBindJson.Status(), errBindJson)
// 		return
// 	}

// 	user := c.MustGet("userData").(entity.User)

// 	newTransactionHistory, err := m.transactionHistoryService.CreateTransactionHistory(user.Id, transactionHistoryRequest)

// 	if err != nil {
// 		c.JSON(err.Status(), err)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, newTransactionHistory)
// }

// func (m transactionHistoryHandler) UpdateTransactionHistoryById(c *gin.Context) {
// 	var transactionHistoryRequest dto.NewTransactionHistoryRequest

// 	if err := c.ShouldBindJSON(&transactionHistoryRequest); err != nil {
// 		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

// 		c.JSON(errBindJson.Status(), errBindJson)
// 		return
// 	}

// 	transactionHistoryId, err := helpers.GetParamId(c, "transactionHistoryId")

// 	if err != nil {
// 		c.AbortWithStatusJSON(err.Status(), err)
// 		return
// 	}

// 	response, err := m.transactionHistoryService.UpdateTransactionHistoryById(transactionHistoryId, transactionHistoryRequest)

// 	if err != nil {
// 		c.AbortWithStatusJSON(err.Status(), err)
// 		return
// 	}

// 	c.JSON(response.StatusCode, response)
// }
