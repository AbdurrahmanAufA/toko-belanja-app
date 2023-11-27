package transactionHistory_repository

import (
	"time"
	"toko-belanja/entity"
)

type product struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type user struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionProduct struct {
	TransactionHistory entity.TransactionHistory
	Product            entity.Product
	User               entity.User
}

type TransactionProductMapped struct {
	Id         int       `json:"Id"`
	UserId     int       `json:"UserId"`
	ProductId  int       `json:"ProductId"`
	Quantity   int       `json:"Quantity"`
	TotalPrice int       `json:"Total_Price"`
	Products   []product `json:"Products"`
	User       []user    `json:"User"`
}

type MyTransactionProduct struct {
	TransactionHistory entity.TransactionHistory
	Product            entity.Product
}

type MyTransactionProductMapped struct {
	Id         int       `json:"Id"`
	UserId     int       `json:"UserId"`
	ProductId  int       `json:"ProductId"`
	Quantity   int       `json:"Quantity"`
	TotalPrice int       `json:"Total_Price"`
	Products   []product `json:"Products"`
}

func (ctm *TransactionProductMapped) HandleMappingTransactionWithProduct(transactionProduct []TransactionProduct) []TransactionProductMapped {
	transactionProductsMapped := make(map[int]TransactionProductMapped)

	for _, eachTransactionProduct := range transactionProduct {
		transactionId := eachTransactionProduct.TransactionHistory.Id
		transactionProductMapped, exists := transactionProductsMapped[transactionId]
		if !exists {
			transactionProductMapped = TransactionProductMapped{
				Id:         eachTransactionProduct.TransactionHistory.Id,
				ProductId:  eachTransactionProduct.TransactionHistory.ProductId,
				UserId:     eachTransactionProduct.TransactionHistory.UserId,
				Quantity:   eachTransactionProduct.TransactionHistory.Quantity,
				TotalPrice: eachTransactionProduct.TransactionHistory.TotalPrice,
			}
		}

		product := product{
			Id:         eachTransactionProduct.Product.Id,
			Title:      eachTransactionProduct.Product.Title,
			Price:      eachTransactionProduct.Product.Price,
			Stock:      eachTransactionProduct.Product.Stock,
			CategoryId: eachTransactionProduct.Product.CategoryId,
			CreatedAt:  eachTransactionProduct.Product.CreatedAt,
			UpdatedAt:  eachTransactionProduct.Product.UpdatedAt,
		}

		user := user{
			Id:        eachTransactionProduct.User.Id,
			Email:     eachTransactionProduct.User.Email,
			FullName:  eachTransactionProduct.User.FullName,
			Balance:   eachTransactionProduct.User.Balance,
			CreatedAt: eachTransactionProduct.User.CreatedAt,
			UpdatedAt: eachTransactionProduct.User.UpdatedAt,
		}

		transactionProductMapped.Products = append(transactionProductMapped.Products, product)
		transactionProductMapped.User = append(transactionProductMapped.User, user)
		transactionProductsMapped[transactionId] = transactionProductMapped
	}

	transactionProducts := []TransactionProductMapped{}
	for _, transactionProduct := range transactionProductsMapped {
		transactionProducts = append(transactionProducts, transactionProduct)
	}
	return transactionProducts
}

func (ctm *MyTransactionProductMapped) HandleMappingMyTransactionWithProduct(mytransactionProduct []MyTransactionProduct) []MyTransactionProductMapped {
	mytransactionProductsMapped := make(map[int]MyTransactionProductMapped)

	for _, eachMyTransactionProduct := range mytransactionProduct {
		mytransactionId := eachMyTransactionProduct.TransactionHistory.Id
		mytransactionProductMapped, exists := mytransactionProductsMapped[mytransactionId]
		if !exists {
			mytransactionProductMapped = MyTransactionProductMapped{
				Id:         eachMyTransactionProduct.TransactionHistory.Id,
				ProductId:  eachMyTransactionProduct.TransactionHistory.ProductId,
				UserId:     eachMyTransactionProduct.TransactionHistory.UserId,
				Quantity:   eachMyTransactionProduct.TransactionHistory.Quantity,
				TotalPrice: eachMyTransactionProduct.TransactionHistory.TotalPrice,
			}
		}

		product := product{
			Id:         eachMyTransactionProduct.Product.Id,
			Title:      eachMyTransactionProduct.Product.Title,
			Price:      eachMyTransactionProduct.Product.Price,
			Stock:      eachMyTransactionProduct.Product.Stock,
			CategoryId: eachMyTransactionProduct.Product.CategoryId,
			CreatedAt:  eachMyTransactionProduct.Product.CreatedAt,
			UpdatedAt:  eachMyTransactionProduct.Product.UpdatedAt,
		}

		mytransactionProductMapped.Products = append(mytransactionProductMapped.Products, product)
		mytransactionProductsMapped[mytransactionId] = mytransactionProductMapped
	}

	mytransactionProducts := []MyTransactionProductMapped{}
	for _, mytransactionProduct := range mytransactionProductsMapped {
		mytransactionProducts = append(mytransactionProducts, mytransactionProduct)
	}
	return mytransactionProducts
}
