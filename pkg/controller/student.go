package controller

import (
	"net/http"
	"strconv"

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
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success", resp)
}

func (sc *StudentController) PostStudent(c echo.Context) error {
	var newStudent domain.Student

	err := c.Bind(&newStudent)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = sc.StudentUsecase.PostStudent(newStudent)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusCreated, "success create student", nil)
}

func (sc *StudentController) UpdateStudent(c echo.Context) error {
	var newStudent domain.Student

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = c.Bind(&newStudent)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = sc.StudentUsecase.UpdateStudent(id, newStudent)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success update student", nil)
}

func (sc *StudentController) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = sc.StudentUsecase.DeleteStudent(id)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success delete student", nil)
}
