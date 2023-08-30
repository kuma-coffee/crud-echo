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

	g.GET("/student", sc.GetStudents)
	g.GET("/student/:id", sc.GetStudentById)
	g.POST("/student", sc.PostStudent)
	g.PUT("/student/:id", sc.UpdateStudent)
	g.DELETE("/student/:id", sc.DeleteStudent)
	g.GET("/search", sc.SearchStudent)
}
