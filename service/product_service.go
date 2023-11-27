package service

import (
	"net/http"
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/pkg/helpers"

	"toko-belanja/repository/product_repository"
)

type ProductService interface {
	CreateProduct(payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	GetProducts() (*dto.GetProductResponse, errs.MessageErr)
	UpdateProductById(productId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	DeleteProduct(productId int) (*dto.DeleteProductResponse, errs.MessageErr)
}

type productService struct {
	productRepo product_repository.Repository
}

func NewProductService(productRepo product_repository.Repository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p *productService) CreateProduct(payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	productRequest := entity.Product{
		Title:      payload.Title,
		Price:      payload.Price,
		Stock:      payload.Stock,
		CategoryId: payload.CategoryId,
	}

	res, err := p.productRepo.CreateProduct(productRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new product data successfully created",
		Data: dto.GetProductForCategory{
			Id:         res.Id,
			Title:      payload.Title,
			Price:      payload.Price,
			Stock:      payload.Stock,
			CategoryId: payload.CategoryId,
			CreatedAt:  res.CreatedAt,
		},
	}

	return &response, err
}

func (ps *productService) GetProducts() (*dto.GetProductResponse, errs.MessageErr) {
	response, err := ps.productRepo.GetProducts()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	return &dto.GetProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Products has been successfully fetched",
		Data:       response,
	}, nil
}

func (p *productService) UpdateProductById(productId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	productRequest := entity.Product{
		Title:      payload.Title,
		Price:      payload.Price,
		Stock:      payload.Stock,
		CategoryId: payload.CategoryId,
	}

	res, err := p.productRepo.UpdateProductById(productId, productRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "product data successfully update",
		Data: dto.GetProductForCategory{
			Id:         res.Id,
			Title:      payload.Title,
			Price:      payload.Price,
			Stock:      payload.Stock,
			CategoryId: payload.CategoryId,
			CreatedAt:  res.CreatedAt,
			UpdatedAt:  res.UpdatedAt,
		},
	}

	return &response, err
}

func (p *productService) DeleteProduct(productId int) (*dto.DeleteProductResponse, errs.MessageErr) {

	err := p.productRepo.DeleteProduct(productId)
	if err != nil {
		return nil, err
	}

	response := dto.DeleteProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Task has been successfully deleted",
	}
	return &response, nil
}