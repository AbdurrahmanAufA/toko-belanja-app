package user_repository

import (
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
)

type Repository interface {
	CreateNewUser(user entity.User) errs.MessageErr
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
	TopUpBalance(payload *entity.User) (*entity.User, errs.MessageErr)
}
