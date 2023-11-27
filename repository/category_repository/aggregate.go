package category_repository

import (
	"toko-belanja/entity"
)

type CategoryWithProduct struct {
	Category entity.Category
	Product  entity.Product
}

type CategoryProductMapped struct {
	Category entity.Category
	Products []entity.Product `json:"products"`
}

func (ctm *CategoryProductMapped) HandleMappingCategoryWithProduct(categoryProduct []CategoryWithProduct) []CategoryProductMapped {
	categoryProductsMapped := []CategoryProductMapped{}

	for _, eachCategoryProduct := range categoryProduct {
		isCategoryExist := false

		for i := range categoryProductsMapped {
			if eachCategoryProduct.Category.Id == categoryProduct[i].Category.Id {
				isCategoryExist = true
				categoryProductsMapped[i].Products = append(categoryProductsMapped[i].Products, eachCategoryProduct.Product)
				break
			}
		}

		if !isCategoryExist {
			categoryProductMapped := CategoryProductMapped{
				Category: eachCategoryProduct.Category,
			}

			categoryProductMapped.Products = append(categoryProductMapped.Products, eachCategoryProduct.Product)
			categoryProductsMapped = append(categoryProductsMapped, categoryProductMapped)
		}
	}
	return categoryProductsMapped
}