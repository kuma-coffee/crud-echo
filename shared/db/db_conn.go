package db

import (
	"database/sql"
	"fmt"

	"github.com/kuma-coffee/crud-echo/config"

	_ "github.com/lib/pq"
)

func NewInstanceDb() *sql.DB {
	conf := config.GetConfig()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBPort, conf.DBUsername, conf.DBPassword, conf.DBName))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
