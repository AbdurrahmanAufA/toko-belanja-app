package service

import (
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/repository/category_repository"
	"toko-belanja/repository/product_repository"
	"toko-belanja/repository/transactionHistory_repository"
	"toko-belanja/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc

	AdminAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo               user_repository.Repository
	categoryRepo           category_repository.Repository
	productRepo            product_repository.Repository
	transactionHistoryRepo transactionHistory_repository.Repository
}

func NewAuthService(userRepo user_repository.Repository, categoryRepo category_repository.Repository, productRepo product_repository.Repository, transactionHistoryRepo transactionHistory_repository.Repository) AuthService {
	return &authService{
		userRepo:               userRepo,
		categoryRepo:           categoryRepo,
		productRepo:            productRepo,
		transactionHistoryRepo: transactionHistoryRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.userRepo.GetUserById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_, err = a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.MustGet("userData").(entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		if user.Role != "admin" {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to access this endpoint, only admin can access it")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
	}
}
