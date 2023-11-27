package dto

import (
	"time"
	"toko-belanja/entity"
)

type NewProductRequest struct {
	Title      string `json:"title" valid:"required~title cannot be empty"`
	Price      int    `json:"price" valid:"required~Price can't be empty, range(0|50000000)~"`
	Stock      int    `json:"stock" valid:"required~Stock can't be empty, range(5|1000000)~"`
	CategoryId int    `json:"category_id"`
}

type NewProductResponse struct {
	Result     string                `json:"result"`
	Message    string                `json:"message"`
	StatusCode int                   `json:"statusCode"`
	Data       GetProductForCategory `json:"data"`
}

type GetProductForCategory struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetProductResponse struct {
	StatusCode int               `json:"code"`
	Message    string            `json:"message"`
	Data       []*entity.Product `json:"data"`
}

type DeleteProductResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
