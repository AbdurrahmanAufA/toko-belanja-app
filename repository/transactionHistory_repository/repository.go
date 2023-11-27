package transactionHistory_repository

import (
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
)

type Repository interface {
	CreateNewTransaction(transactionPayLoad *entity.TransactionHistory) (*dto.TransactionBill, errs.MessageErr)
	GetMyTransaction(UserId int) ([]MyTransactionProductMapped, errs.MessageErr)
	GetTransaction() ([]TransactionProductMapped, errs.MessageErr)
	// UpdateTransactionHistoryById(payload entity.TransactionHistory) errs.MessageErr
}
