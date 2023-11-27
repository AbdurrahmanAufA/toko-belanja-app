package dto

import (
	"time"
	"toko-belanja/entity"
)

type NewTransactionRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity" valid:"required~quantity cannot be empty"`
}

type NewTransactionResponse struct {
	Result     string            `json:"result"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Data       TransactionReturn `json:"data"`
}

type TransactionReturn struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionRequest struct {
	ProductId int  `json:"product_id" example:"1"`
	Quantity  int `json:"quantity" valid:"required~Quantity can't be empty"`
}

type TransactionBill struct {
	TotalPrice   int   `json:"total_price"`
	Quantity     int   `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionResponse struct {
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Balance   int      `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MyTransaction struct {
	Id         int            `json:"id"`
	UserId     int            `json:"user_id"`
	ProductId  int            `json:"product_id"`
	Quantity   int           `json:"quantity"`
	TotalPrice int           `json:"total_price"`
	Products   entity.Product `json:"products"`
}

type UsersTransaction struct {
	Id         int            `json:"id"`
	UserId     int            `json:"user_id"`
	ProductId  int            `json:"product_id"`
	Quantity   int           `json:"quantity"`
	TotalPrice int           `json:"total_price"`
	Products   entity.Product `json:"products"`
	User       User           `json:"users"`
}

type TransactionHistoryResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
