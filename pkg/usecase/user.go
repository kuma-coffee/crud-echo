package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kuma-coffee/crud-echo/pkg/domain"
	"github.com/kuma-coffee/crud-echo/pkg/dto"
	"github.com/kuma-coffee/crud-echo/pkg/helper"
	"github.com/mitchellh/mapstructure"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepository}
}

func (uu UserUsecase) Login(loginDTO dto.LoginDTO) (string, error) {
	dbUser, err := uu.UserRepository.GetUserEmail(loginDTO.Email)
	if err != nil {
		return fmt.Sprint(), err
	}

	if dbUser.Email == "" || dbUser.Id == 0 {
		return fmt.Sprint(), errors.New("user not found")
	}

	match, err := helper.CheckPasswordHash(loginDTO.Password, dbUser.Password)
	if !match {
		return fmt.Sprint(), err
	}

	claims := &domain.JwtCustomClaims{
		Email: loginDTO.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func (uu UserUsecase) Register(userDTO dto.UserDTO) error {
	var user domain.User
	mapstructure.Decode(userDTO, &user)

	dbUser, _ := uu.UserRepository.GetUserEmail(user.Email)
	if dbUser.Email != "" {
		return errors.New("email already registered")
	}

	passwordHash, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = passwordHash

	return uu.UserRepository.Register(user)
}

func (uu UserUsecase) GetUserEmail(email string) (domain.User, error) {
	return uu.UserRepository.GetUserEmail(email)
}
