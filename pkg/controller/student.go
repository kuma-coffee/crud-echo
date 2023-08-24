package controller

import (
	"net/http"

	"github.com/kuma-coffee/crud-echo/pkg/domain"
	"github.com/kuma-coffee/crud-echo/shared/response"
	"github.com/labstack/echo/v4"
)

type StudentController struct {
	StudentUsecase domain.StudentUsecase
}

func (sc *StudentController) GetStudents(c echo.Context) error {
	resp, err := sc.StudentUsecase.GetStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return response.SetResponse(c, http.StatusOK, "success", resp)
}
