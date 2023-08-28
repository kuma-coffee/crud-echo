package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kuma-coffee/crud-echo/pkg/dto"
)

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type UserRepository interface {
	Register(user User) error
	GetUserEmail(email string) (User, error)
}

type UserUsecase interface {
	Login(loginDTO dto.LoginDTO) (string, error)
	Register(userDTO dto.UserDTO) error
	GetUserEmail(email string) (User, error)
}
