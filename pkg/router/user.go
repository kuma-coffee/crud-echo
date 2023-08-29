package router

import (
	"database/sql"

	"github.com/kuma-coffee/crud-echo/pkg/controller"
	"github.com/kuma-coffee/crud-echo/pkg/repository"
	"github.com/kuma-coffee/crud-echo/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := &controller.UserController{
		UserUsecase: uu,
	}

	g.POST("/login", uc.Login)
	g.POST("/register", uc.Register)
}
