package dto

import (
	"time"
)

type NewUserRequest struct {
	FullName string `json:"full_name" valid:"required~full_name cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty, length(6|255)~Minimum password is 6 length"`
}

type NewUserResponse struct {
	Result     string   `json:"result"`
	StatusCode int      `json:"statusCode"`
	Message    string   `json:"message"`
	Data       Response `json:"data"`
}

type AdminUserResponse struct {
	Result     string   `json:"result"`
	StatusCode int      `json:"statusCode"`
	Message    string   `json:"message"`
	Data       Response `json:"data"`
}

type UserReturn struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Response struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UserRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty, length(6|255)~Minimum password is 6 length"`
}

type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}

type TopupRequest struct {
	Balance int `json:"balance" valid:"required~Balance can't be empty, range(0|100000000)~Balance can't be less than 0 or more than 100000000" example:"150000"`
}

type TopupResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
