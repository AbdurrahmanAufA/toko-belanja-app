package database

import (
	"database/sql"
	"fmt"
	"log"
	"toko-belanja/infra/config"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.GetAppConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {

	usersTable := `
		CREATE TABLE IF NOT EXISTS "users" (
			id SERIAL PRIMARY KEY,
			full_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password TEXT NOT NULL,
			role VARCHAR(255) NOT NULL,
			balance int NOT NULL DEFAULT 0,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now(),
			CONSTRAINT
				unique_email
					UNIQUE(email)
		);
	`
	categoryTable := `
		CREATE TABLE IF NOT EXISTS "category" (
			id SERIAL PRIMARY KEY,
			type VARCHAR(255) NOT NULL,
			sold_product_amount int NOT NULL DEFAULT 0,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now()
		);
	`

	productTable := `
		CREATE TABLE IF NOT EXISTS "product" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			price int NOT NULL,
			stock int NOT NULL,
			category_id int NOT NULL,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now(),
			CONSTRAINT product_category_id_fk
				FOREIGN KEY(category_id)
					REFERENCES category(id)
		);
	`

	transactionTable := `
		CREATE TABLE IF NOT EXISTS "transaction_history" (
			id SERIAL PRIMARY KEY,
			product_id int NOT NULL,
			user_id int NOT NULL,
			quantity int NOT NULL,
			total_price int NOT NULL,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now(),
			CONSTRAINT transaction_product_id_fk
				FOREIGN KEY(product_id)
					REFERENCES product(id),
			CONSTRAINT transaction_user_id_fk 
				FOREIGN KEY(user_id) 
					REFERENCES users(id) 
		);
	`
	createTrigger := `
			CREATE OR REPLACE FUNCTION reduce_balance_on_transaction() RETURNS TRIGGER AS $$
			BEGIN
				UPDATE users
				SET balance = balance - (NEW.quantity * (
					SELECT price FROM product WHERE id = NEW.product_id
				))
				WHERE id = NEW.user_id;
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			CREATE OR REPLACE TRIGGER transaction_balance_trigger
			AFTER INSERT ON transaction_history
			FOR EACH ROW
			EXECUTE FUNCTION reduce_balance_on_transaction();


			CREATE OR REPLACE FUNCTION reduce_stock_on_transaction() RETURNS TRIGGER AS $$
			BEGIN
				UPDATE product
				SET stock = stock - NEW.quantity
				WHERE id = NEW.product_id;
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			CREATE OR REPLACE TRIGGER transaction_stock_trigger
			AFTER INSERT ON transaction_history
			FOR EACH ROW
			EXECUTE FUNCTION reduce_stock_on_transaction();

			CREATE OR REPLACE FUNCTION increase_sold_amount_on_transaction() RETURNS TRIGGER AS $$
			BEGIN
				UPDATE category
				SET sold_product_amount = sold_product_amount + NEW.quantity
				WHERE id = (SELECT category_id FROM product WHERE id = NEW.product_id);
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			CREATE OR REPLACE TRIGGER transaction_category_trigger
			AFTER INSERT ON transaction_history
			FOR EACH ROW
			EXECUTE FUNCTION increase_sold_amount_on_transaction();
		`

	createTableQueries := fmt.Sprintf("%s %s %s %s", usersTable, categoryTable, productTable, transactionTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}

	_, err = db.Exec(createTrigger)

	if err != nil {
		log.Panic("error while create OR REPLACE trigger: ", err.Error())
		return
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
