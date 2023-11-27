package service

import (
	"fmt"
	"net/http"
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/pkg/helpers"
	"toko-belanja/repository/product_repository"
	"toko-belanja/repository/transactionHistory_repository"
	"toko-belanja/repository/user_repository"
)

type TransactionHistoryService interface {
	// CreateTransaction(newTransactionRequest dto.NewTransactionRequest, userId int) (*dto.TransactionResponse, errs.MessageErr)
	CreateTransaction(productId int, transactionPayLoad *dto.TransactionRequest) (*dto.TransactionHistoryResponse, errs.MessageErr)
	GetTransactionWithProducts(userId int) (*dto.TransactionHistoryResponse, errs.MessageErr)
	GetTransactionWithProductsAndUser() (*dto.TransactionHistoryResponse, errs.MessageErr)
}

type transactionHistoryService struct {
	transactionHistoryRepo transactionHistory_repository.Repository
	productRepo            product_repository.Repository
	userRepo               user_repository.Repository
}

func NewTransactionHistoryService(transactionHistoryRepo transactionHistory_repository.Repository, productRepo product_repository.Repository, userRepo user_repository.Repository) TransactionHistoryService {
	return &transactionHistoryService{
		transactionHistoryRepo: transactionHistoryRepo,
		productRepo:            productRepo,
		userRepo:               userRepo,
	}
}

func (ts *transactionHistoryService) CreateTransaction(userId int, transactionPayLoad *dto.TransactionRequest) (*dto.TransactionHistoryResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(transactionPayLoad)

	if err != nil {
		return nil, err
	}

	product, err := ts.productRepo.GetProductById(transactionPayLoad.ProductId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewNotFoundError("Product not Found")
		}
		return nil, err
	}

	if transactionPayLoad.Quantity > product.Stock {
		return nil, errs.NewUnprocessibleEntityError("Insufficient stock for the requested quantity")
	}

	user, err := ts.userRepo.GetUserById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("Not Found")
		}
		return nil, err
	}

	totalCost := product.Price * transactionPayLoad.Quantity
	if user.Balance < totalCost {
		fmt.Println(totalCost)
		fmt.Println(user.Balance)
		return nil, errs.NewUnprocessibleEntityError("Insufficient balance for the transaction")
	}

	transaction := &entity.TransactionHistory{
		UserId:    userId,
		ProductId: transactionPayLoad.ProductId,
		Quantity:  transactionPayLoad.Quantity,
	}

	response, err := ts.transactionHistoryRepo.CreateNewTransaction(transaction)

	if err != nil {
		return nil, err
	}

	return &dto.TransactionHistoryResponse{
		Code:    http.StatusCreated,
		Message: "You have successfully purchased the product",
		Data: dto.TransactionBill{
			TotalPrice:   response.TotalPrice,
			Quantity:     response.Quantity,
			ProductTitle: response.ProductTitle,
		},
	}, nil
}

func (ts *transactionHistoryService) GetTransactionWithProducts(userId int) (*dto.TransactionHistoryResponse, errs.MessageErr) {
	response, err := ts.transactionHistoryRepo.GetMyTransaction(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("Not Found")
		}
		return nil, err
	}

	return &dto.TransactionHistoryResponse{
		Code:    http.StatusOK,
		Message: "Your transaction has been successfully fetched",
		Data:    response,
	}, nil
}

func (ts *transactionHistoryService) GetTransactionWithProductsAndUser() (*dto.TransactionHistoryResponse, errs.MessageErr) {
	response, err := ts.transactionHistoryRepo.GetTransaction()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	return &dto.TransactionHistoryResponse{
		Code:    http.StatusOK,
		Message: "Transaction has been successfuly fetched",
		Data:    response,
	}, nil
}
