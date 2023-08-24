package cmd

import (
	"github.com/kuma-coffee/crud-echo/pkg/router"
	"github.com/kuma-coffee/crud-echo/shared/db"
	"github.com/labstack/echo/v4"
)

func RunServer() {
	e := echo.New()
	g := e.Group("")

	Apply(e, g)

	e.Logger.Error(e.Start(":5000"))
}

func Apply(e *echo.Echo, g *echo.Group) {
	db := db.NewInstanceDb()
	router.NewStudentRouter(e, g, db)
}
