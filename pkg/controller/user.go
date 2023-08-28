package controller

import (
	"net/http"

	"github.com/kuma-coffee/crud-echo/pkg/domain"
	"github.com/kuma-coffee/crud-echo/pkg/dto"
	"github.com/kuma-coffee/crud-echo/shared/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) Login(c echo.Context) error {
	var loginDTO dto.LoginDTO

	err := c.Bind(&loginDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = loginDTO.ValidationLogin()
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	res, err := uc.UserUsecase.Login(loginDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusCreated, "user login success", res)
}

func (uc *UserController) Register(c echo.Context) error {
	var userDTO dto.UserDTO

	err := c.Bind(&userDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = userDTO.ValidationUser()
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = uc.UserUsecase.Register(userDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusCreated, "register user success", nil)
}
