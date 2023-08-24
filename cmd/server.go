package cmd

import (
	"github.com/kuma-coffee/crud-echo/shared/db"
	"github.com/labstack/echo/v4"
)

func RunServer() {
	e := echo.New()

	db.NewInstanceDb()

	e.Logger.Error(e.Start(":5000"))
}
