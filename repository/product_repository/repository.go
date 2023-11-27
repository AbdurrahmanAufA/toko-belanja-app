package product_repository

import (
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
)

type Repository interface {
	CreateProduct(productPayload entity.Product) (*entity.Product, errs.MessageErr)
	GetProducts() ([]*entity.Product, errs.MessageErr)
	GetProductById(productId int) (*entity.Product, errs.MessageErr)
	UpdateProductById(productId int, payload entity.Product) (*entity.Product, errs.MessageErr)
	DeleteProduct(productId int) errs.MessageErr
}
