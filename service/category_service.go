package service

import (
	"net/http"
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/pkg/helpers"
	"toko-belanja/repository/category_repository"
)

type CategoryService interface {
	CreateNewCategory(userId int, payload dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	GetCategoryById(userId int) (*dto.GetCategoryResponse, errs.MessageErr)
	PatchCategory(categoryId int, payload dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(categoryId int) (*dto.DeleteCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repository.Repository
}

func NewCategoryService(categoryRepo category_repository.Repository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c *categoryService) CreateNewCategory(userId int, payload dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr) {

	categoryRequest := &entity.Category{
		Type:              payload.Type,
		SoldProductAmount: 0,
	}

	ctg, err := c.categoryRepo.CreateNewCategory(categoryRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewCategoryResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new category data successfully created",
		Data: dto.CategoryResponse{
			Id:                ctg.Id,
			Type:              categoryRequest.Type,
			SoldProductAmount: categoryRequest.SoldProductAmount,
			CreatedAt:         ctg.CreatedAt,
		},
	}

	return &response, err
}

func (c *categoryService) GetCategoryById(userId int) (*dto.GetCategoryResponse, errs.MessageErr) {

	categories, err := c.categoryRepo.GetCategory()

	if err != nil {
		return nil, err
	}

	categoryResult := []dto.GetCategoryReturn{}

	for _, eachCategory := range categories {
		category := dto.GetCategoryReturn{
			Id:        eachCategory.Category.Id,
			Type:      eachCategory.Category.Type,
			CreatedAt: eachCategory.Category.CreatedAt,
			UpdatedAt: eachCategory.Category.UpdatedAt,
			Product:   []dto.GetProductForCategory{},
		}

		for _, eachProduct := range eachCategory.Products {
			product := dto.GetProductForCategory{
				Id:         eachProduct.Id,
				Title:      eachProduct.Title,
				Price:      eachProduct.Price,
				Stock:      eachProduct.Stock,
				CategoryId: eachProduct.CategoryId,
				CreatedAt:  eachProduct.CreatedAt,
				UpdatedAt:  eachProduct.UpdatedAt,
			}
			category.Product = append(category.Product, product)
		}

		categoryResult = append(categoryResult, category)
	}

	response := dto.GetCategoryResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "Get category data success",
		Data:       categoryResult,
	}
	return &response, nil
}

func (c *categoryService) PatchCategory(categoryId int, payload dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	categoryUpdate := entity.Category{
		Type: payload.Type,
	}

	updateCategory, err := c.categoryRepo.UpdateCategory(categoryId, categoryUpdate)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateCategoryResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "category successfully updated",
		Data: dto.UpdateCategoryReturn{
			Id:        updateCategory.Id,
			Type:      payload.Type,
			UpdatedAt: updateCategory.UpdatedAt,
		},
	}

	return &response, nil
}

func (c *categoryService) DeleteCategory(categoryId int) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	_, err := c.categoryRepo.GetCategoryById(categoryId)
	if err != nil {
		return nil, err
	}

	err = c.categoryRepo.DeleteCategory(categoryId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteCategoryResponse{
		StatusCode: http.StatusOK,
		Message:    "Category has been successfully deleted",
	}

	return &response, nil
}