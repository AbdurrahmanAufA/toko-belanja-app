package user_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/repository/user_repository"
)

const (
	retrieveUserByEmail = `
		SELECT id, full_name, email, password, role
		FROM "users"
		WHERE email = $1;
	`

	retrieveUserById = `
		SELECT id, full_name, email, password, role, balance
		FROM "users"
		WHERE id = $1;
	`

	createNewUser = `
		INSERT INTO "users"
		(
			full_name,
			email,
			password,
			role,
			balance

		)
		VALUES ($1, $2, $3, $4, $5)
	`

	usersTopUp = `
		UPDATE "users"
		SET balance = balance + $2
		WHERE id = $1
		RETURNING
			balance
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.Repository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(user entity.User) errs.MessageErr {
	_, err := u.db.Exec(createNewUser, user.FullName, user.Email, user.Password, user.Role, user.Balance)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

// func (u *userPG) CreateNewUser(user entity.User) (*entity.User, errs.MessageErr) {

// 	var newUser entity.User

// 	rows := u.db.QueryRow(createNewUser, user.FullName, user.Email, user.Password, user.Role, user.Balance)

// 	err := rows.Scan(&newUser.Id, &newUser.CreatedAt)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errs.NewNotFoundError("user not found")
// 		}

// 		return nil, errs.NewInternalServerError("Something went wrong")
// 	}

// 	return &user, nil
// }

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserById, userId)

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Role, &user.Balance)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserByEmail, userEmail)

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) TopUpBalance(payload *entity.User) (*entity.User, errs.MessageErr) {

	_, err := u.db.Exec(usersTopUp, payload.Id, payload.Balance)

	if err != nil {
		fmt.Println(err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return payload, nil
}
