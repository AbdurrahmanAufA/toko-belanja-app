package transactionHistory_pg

import (
	"database/sql"
	"fmt"
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/repository/transactionHistory_repository"
)

const (
	createTransaction = `
	INSERT INTO transaction_history (user_id, product_id, quantity, total_price)
	VALUES ($1, $2, $3, ((SELECT p.price
	FROM product AS p
	WHERE
	id =$2)*$3))
	RETURNING
	product_id, quantity, (
		SELECT 
			title 
		FROM 
			product
		WHERE 
			id = $2
	)
	`

	getMyTransaction = `
		SELECT t.id, 
			t.product_id,
			t.user_id, 
			t.quantity, 
			t.total_price, 
			p.id, 
			p.title, 
			p.price, 
			p.stock, 
			p.category_Id, 
			p.createdat, 
			p.updatedat
		FROM 
			transaction_history as t
		LEFT JOIN
			product as p
		ON
			t.product_id = p.id
		WHERE t.user_id = $1
		ORDER BY 
		t.id ASC
		`
	getTransaction = `
		SELECT
			t.id,
			t.product_id,
			t.user_id,
			t.quantity,
			t.total_price,
			p.id AS product_id,
			p.title AS product_title,
			p.price,
			p.stock,
			p.category_id,
			p.createdat AS product_created_at,
			p.updatedat AS product_updated_at,
			u.id AS user_id,
			u.email,
			u.full_name,
			u.balance,
			u.createdat AS user_created_at,
			u.updatedat AS user_updated_at
		FROM
			transaction_history AS t
		LEFT JOIN
			product AS p
		ON
			t.product_id = p.id
		LEFT JOIN
			users AS u
		ON
			t.user_id = u.id
		ORDER BY
			t.id ASC;
	`
)

type transactionHistoryPG struct {
	db *sql.DB
}

func NewTransactionHistoryPG(db *sql.DB) transactionHistory_repository.Repository {
	return &transactionHistoryPG{
		db: db,
	}
}

func (t *transactionHistoryPG) CreateNewTransaction(transactionPayLoad *entity.TransactionHistory) (*dto.TransactionBill, errs.MessageErr) {
	tx, err := t.db.Begin()
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	transaction := dto.TransactionBill{}

	row := tx.QueryRow(
		createTransaction,
		transactionPayLoad.UserId,
		transactionPayLoad.ProductId,
		transactionPayLoad.Quantity,
	)
	err = row.Scan(
		&transaction.TotalPrice,
		&transaction.Quantity,
		&transaction.ProductTitle,
	)

	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &transaction, nil
}

func (t *transactionHistoryPG) GetMyTransaction(UserId int) ([]transactionHistory_repository.MyTransactionProductMapped, errs.MessageErr) {
	mytransactionProducts := []transactionHistory_repository.MyTransactionProduct{}
	rows, err := t.db.Query(getMyTransaction, UserId)

	if err != nil {
		fmt.Println(err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		var mytransactionProduct transactionHistory_repository.MyTransactionProduct

		err := rows.Scan(
			&mytransactionProduct.TransactionHistory.Id,
			&mytransactionProduct.TransactionHistory.ProductId,
			&mytransactionProduct.TransactionHistory.UserId,
			&mytransactionProduct.TransactionHistory.Quantity,
			&mytransactionProduct.TransactionHistory.TotalPrice,
			&mytransactionProduct.Product.Id,
			&mytransactionProduct.Product.Title,
			&mytransactionProduct.Product.Price,
			&mytransactionProduct.Product.Stock,
			&mytransactionProduct.Product.CategoryId,
			&mytransactionProduct.Product.CreatedAt,
			&mytransactionProduct.Product.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		mytransactionProducts = append(mytransactionProducts, mytransactionProduct)
	}

	result := transactionHistory_repository.MyTransactionProductMapped{}
	return result.HandleMappingMyTransactionWithProduct(mytransactionProducts), nil
}

func (t *transactionHistoryPG) GetTransaction() ([]transactionHistory_repository.TransactionProductMapped, errs.MessageErr) {
	transactionProducts := []transactionHistory_repository.TransactionProduct{}
	rows, err := t.db.Query(getTransaction)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		var transactionProduct transactionHistory_repository.TransactionProduct

		err := rows.Scan(
			&transactionProduct.TransactionHistory.Id,
			&transactionProduct.TransactionHistory.ProductId,
			&transactionProduct.TransactionHistory.UserId,
			&transactionProduct.TransactionHistory.Quantity,
			&transactionProduct.TransactionHistory.TotalPrice,
			&transactionProduct.Product.Id,
			&transactionProduct.Product.Title,
			&transactionProduct.Product.Price,
			&transactionProduct.Product.Stock,
			&transactionProduct.Product.CategoryId,
			&transactionProduct.Product.CreatedAt,
			&transactionProduct.Product.UpdatedAt,
			&transactionProduct.User.Id,
			&transactionProduct.User.Email,
			&transactionProduct.User.FullName,
			&transactionProduct.User.Balance,
			&transactionProduct.User.CreatedAt,
			&transactionProduct.User.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		transactionProducts = append(transactionProducts, transactionProduct)
	}

	result := transactionHistory_repository.TransactionProductMapped{}
	return result.HandleMappingTransactionWithProduct(transactionProducts), nil
}
