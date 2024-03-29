package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/kuma-coffee/crud-echo/pkg/domain"
	"github.com/kuma-coffee/crud-echo/pkg/dto"
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

	return response.SetResponse(c, http.StatusOK, "success get students", resp)
}

func (sc *StudentController) GetStudentById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := sc.StudentUsecase.GetStudentById(id)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success get student", resp)
}

func (sc *StudentController) PostStudent(c echo.Context) error {
	var studentDTO dto.StudentDTO

	err := c.Bind(&studentDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = studentDTO.ValidationStudent()
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = sc.StudentUsecase.PostStudent(studentDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusCreated, "success create student", nil)
}

func (sc *StudentController) UpdateStudent(c echo.Context) error {
	var studentDTO dto.StudentDTO

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = c.Bind(&studentDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	_, err = sc.StudentUsecase.GetStudentById(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "student id not found", nil)
	}

	err = studentDTO.ValidationStudent()
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = sc.StudentUsecase.UpdateStudent(id, studentDTO)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success update student", nil)
}

func (sc *StudentController) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	_, err = sc.StudentUsecase.GetStudentById(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "student id not found", nil)
	}

	err = sc.StudentUsecase.DeleteStudent(id)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success delete student", nil)
}

func (sc *StudentController) SearchStudent(c echo.Context) error {
	query := c.QueryParam("q")
	querySplit := strings.Fields(query)

	resp, err := sc.StudentUsecase.SearchStudent(querySplit)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success search student", resp)
}

func (sc *StudentController) SortStudent(c echo.Context) error {
	query := c.QueryParam("q")

	resp, err := sc.StudentUsecase.SortStudent(query)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success sort student", resp)
}
