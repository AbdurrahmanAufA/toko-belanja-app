package service

import (
	"fmt"
	"net/http"
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/pkg/helpers"
	"toko-belanja/repository/user_repository"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(userRequest dto.UserRequest) (*dto.LoginResponse, errs.MessageErr)
	TopUpBalance(userId int, userRequest *dto.TopupRequest) (*dto.TopupResponse, errs.MessageErr)
	SeedAdminUser() (*dto.NewUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.Repository
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Login(userRequest dto.UserRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userRequest)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(userRequest.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(userRequest.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully logged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}

	return &response, nil
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	existingEmail, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingEmail != nil {
		return nil, errs.NewDuplicateDataError("Please Try Another Email")
	}

	user := entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     "customer",
		Balance:  0,
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(user)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "user registered successfully",
		Data: dto.Response{
			Id:        user.Id,
			FullName:  user.FullName,
			Email:     user.Email,
			Password:  user.Password,
			Balance:   user.Balance,
			CreatedAt: user.CreatedAt,
		},
	}

	return &response, nil
}

func (us *userService) TopUpBalance(userId int, userRequest *dto.TopupRequest) (*dto.TopupResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userRequest)

	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Id:      userId,
		Balance: userRequest.Balance,
	}

	res, err := us.userRepo.TopUpBalance(payload)

	if err != nil {
		return nil, err
	}

	massage := fmt.Sprintf("Your balance has been successfully updated to Rp%d", res.Balance)
	response := dto.TopupResponse{
		StatusCode: http.StatusOK,
		Message:    massage,
	}
	return &response, nil
}

func (u *userService) SeedAdminUser() (*dto.NewUserResponse, errs.MessageErr) {

	existingAdmin, err := u.userRepo.GetUserByEmail("admin@gmail.com")
	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingAdmin != nil {
		fmt.Println("Admin user already exists. Skipping admin seeding.")
		return nil, nil
	}

	adminUser := entity.User{
		FullName: "Admin Control",
		Email:    "admin@gmail.com",
		Password: "123456",
		Balance:  0,
		Role:     "admin",
	}

	err = adminUser.HashPassword()
	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(adminUser)
	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Admin user registered successfully",
		Data: dto.Response{
			Id:        adminUser.Id,
			FullName:  adminUser.FullName,
			Email:     adminUser.Password,
			Password:  adminUser.Password,
			Balance:   adminUser.Balance,
			CreatedAt: adminUser.CreatedAt,
		},
	}

	return &response, nil
}
