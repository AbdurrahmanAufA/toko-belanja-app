package category_repository

import (
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
)

type Repository interface {
	CreateNewCategory(newCategory *entity.Category) (*entity.Category, errs.MessageErr)
	GetCategoryById(categoryId int) (*entity.Category, errs.MessageErr)
	GetCategory() ([]CategoryProductMapped, errs.MessageErr)
	UpdateCategory(categoryId int, update entity.Category) (*entity.Category, errs.MessageErr)
	DeleteCategory(categoryId int) errs.MessageErr
}
