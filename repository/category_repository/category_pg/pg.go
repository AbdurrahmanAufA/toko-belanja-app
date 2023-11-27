package category_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/repository/category_repository"
)

const (
	createCategory = `
		INSERT INTO "category"
		(
			type,
			sold_product_amount
		)
		VALUES ($1,$2)
		RETURNING id,createdat
	`

	getCategoryByIdQuery = `
		SELECT id, type, sold_product_amount, createdat, updatedat from "category"
		WHERE id = $1;
	`

	GetCategoryWithProduct = `
		SELECT "c"."id", "c"."type", "c"."sold_product_amount", "c"."createdat", "c"."updatedat", "t"."id", "t"."title", "t"."price", "t"."stock", "t"."category_id", "t"."createdat", "t"."updatedat"
		FROM "category" as "c"
		LEFT JOIN "product" as "t" ON "c"."id" = "t"."category_id"
	`

	updateCategoryByIdQuery = `
		UPDATE "category"
		SET type = $2
		WHERE id = $1
		RETURNING id,updatedat
	`

	deleteCategoryById = `
		DELETE FROM "category"
		WHERE id = $1
	`

)

type categoryPG struct {
	db *sql.DB
}

func NewCategoryPG(db *sql.DB) category_repository.Repository {
	return &categoryPG{
		db: db,
	}
}

func (m *categoryPG) GetCategoryById(categoryId int) (*entity.Category, errs.MessageErr) {
	row := m.db.QueryRow(getCategoryByIdQuery, categoryId)

	var category entity.Category

	err := row.Scan(&category.Id, &category.Type, &category.SoldProductAmount, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("category not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) CreateNewCategory(newCategory *entity.Category) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	rows := c.db.QueryRow(createCategory, newCategory.Type, newCategory.SoldProductAmount)

	err := rows.Scan(&category.Id, &category.CreatedAt)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &category, nil
}

func (c *categoryPG) GetCategory() ([]category_repository.CategoryProductMapped, errs.MessageErr) {
	rows, err := c.db.Query(GetCategoryWithProduct)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return nil, errs.NewInternalServerError("something went wrong")
	}
	categoryProducts := []category_repository.CategoryWithProduct{}

	for rows.Next() {
		var categoryProduct category_repository.CategoryWithProduct

		err = rows.Scan(
			&categoryProduct.Category.Id, &categoryProduct.Category.Type, &categoryProduct.Category.SoldProductAmount, &categoryProduct.Category.CreatedAt, &categoryProduct.Category.UpdatedAt,
			&categoryProduct.Product.Id, &categoryProduct.Product.Title, &categoryProduct.Product.Price, &categoryProduct.Product.Stock, &categoryProduct.Product.CategoryId, &categoryProduct.Product.CreatedAt, &categoryProduct.Product.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError("Please fill all category with troduct to get all category data")
		}
		categoryProducts = append(categoryProducts, categoryProduct)
	}
	var result category_repository.CategoryProductMapped

	return result.HandleMappingCategoryWithProduct(categoryProducts), nil
}

func (c *categoryPG) UpdateCategory(categoryId int, update entity.Category) (*entity.Category, errs.MessageErr) {
	var category entity.Category
	rows := c.db.QueryRow(updateCategoryByIdQuery, categoryId, update.Type)

	err := rows.Scan(&category.Id, &category.UpdatedAt)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) DeleteCategory(categoryId int) errs.MessageErr {
	_, err := c.db.Exec(deleteCategoryById, categoryId)
	fmt.Println(err)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}