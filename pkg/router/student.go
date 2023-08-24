package router

import (
	"database/sql"

	"github.com/kuma-coffee/crud-echo/pkg/controller"
	"github.com/kuma-coffee/crud-echo/pkg/repository"
	"github.com/kuma-coffee/crud-echo/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func NewStudentRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	sr := repository.NewStudentRepository(db)
	su := usecase.NewStudentUsecase(sr)
	sc := &controller.StudentController{
		StudentUsecase: su,
	}

	e.GET("/student", sc.GetStudents)
	e.POST("/student", sc.PostStudent)
	e.PUT("/student/:id", sc.UpdateStudent)
	e.DELETE("/student/:id", sc.DeleteStudent)
}
