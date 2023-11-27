package dto

import "time"

type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty" example:"Jelangkung"`
}

type NewCategoryResponse struct {
	Result     string           `json:"result"`
	Message    string           `json:"message"`
	StatusCode int              `json:"statusCode"`
	Data       CategoryResponse `json:"data"`
}

type CategoryResponse struct {
	Id                int       `json:"id" example:"1"`
	Type              string    `json:"type" example:"jersey"`
	SoldProductAmount int       `json:"sold_product_amount" example:"0"`
	CreatedAt         time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type GetCategoryResponse struct {
	Result     string              `json:"result"`
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	Data       []GetCategoryReturn `json:"data"`
}

type GetCategoryReturn struct {
	Id                int                     `json:"id"`
	Type              string                  `json:"type"`
	SoldProductAmount int                     `json:"sold_product_amount" example:"0"`
	CreatedAt         time.Time               `json:"created_at"`
	UpdatedAt         time.Time               `json:"Updated_at"`
	Product           []GetProductForCategory `json:"Tasks"`
}

type UpdateCategoryReturn struct {
	Id         int       `json:"id"`
	Type       string    `json:"type"`
	SoldProductAmount int `json:"sold_product_amount"`
	UpdatedAt time.Time `json:"Updated_at"`
}
type UpdateCategoryResponse struct {
	Result     string               `json:"result"`
	StatusCode int                  `json:"statusCode"`
	Message    string               `json:"message"`
	Data       UpdateCategoryReturn `json:"data"`
}

type DeleteCategoryResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}