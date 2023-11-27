package product_pg

import (
	"database/sql"
	"fmt"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/repository/product_repository"
)

const (
	checkCategoryExistence = `
		SELECT id FROM "category" WHERE id = $1
	`

	createProduct = `
		INSERT INTO "product"
		(
			title,
			price,
			stock,
			category_id
		)
		VALUES ($1,$2,$3,$4)
		RETURNING id,createdat
	`
	getProduct = `
		SELECT id, title, price, stock, category_id, createdat FROM "product"
	`

	getProductById = `
		SELECT id, title, price, stock, category_id, createdat,updatedat
		FROM "product"
		WHERE id = $1
	`

	updateProductById = `
		UPDATE "product"
		SET title = $2,
			price = $3,
			stock = $4,
			category_Id = $5
		WHERE id = $1
		RETURNING id, title, price, stock, category_id, createdat, updatedat
	`

	deleteProductId = `
		DELETE FROM "product"
		WHERE id = $1
	`
)

type productPG struct {
	db *sql.DB
}

func NewProductPG(db *sql.DB) product_repository.Repository {
	return &productPG{
		db: db,
	}
}

func (p *productPG) CreateProduct(payload entity.Product) (*entity.Product, errs.MessageErr) {
	var categoryRow int
	var product entity.Product

	err := p.db.QueryRow(checkCategoryExistence, payload.CategoryId).Scan(&categoryRow)

	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("product not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	err = p.db.QueryRow(createProduct, payload.Title, payload.Price, payload.Stock, payload.CategoryId).Scan(&product.Id, &product.CreatedAt)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("product not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &product, nil
}

func (p *productPG) GetProducts() ([]*entity.Product, errs.MessageErr) {
	rows, err := p.db.Query(getProduct)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	products := []*entity.Product{}

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &product.CreatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("Please fill all category with troduct to get all category data")
		}

		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return products, nil
}

func (p *productPG) UpdateProductById(productId int, payload entity.Product) (*entity.Product, errs.MessageErr) {
	var categoryRow int
	var product entity.Product

	err := p.db.QueryRow(checkCategoryExistence, payload.CategoryId).Scan(&categoryRow)

	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	err = p.db.QueryRow(updateProductById, productId, payload.Title, payload.Price, payload.Stock, payload.CategoryId).Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("product not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &product, nil
}

func (p *productPG) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	var product entity.Product

	rows := p.db.QueryRow(getProductById, productId)

	err := rows.Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &product, nil
}

func (p *productPG) DeleteProduct(productId int) errs.MessageErr {

	_, err := p.db.Exec(deleteProductId, productId)
	fmt.Println(err)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}
