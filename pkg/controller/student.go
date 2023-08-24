package controller

import (
	"net/http"

	"github.com/kuma-coffee/crud-echo/pkg/domain"
	"github.com/labstack/echo/v4"
)

type StudentController struct {
	StudentUsecase domain.StudentUsecase
}

func (sc *StudentController) GetStudents(c echo.Context) error {
	students, err := sc.StudentUsecase.GetStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, students)
}
