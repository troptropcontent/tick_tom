package db_initializer

import (
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/env"
)

func Init() {
	db.DB = db.New(db.Config{
		Host:     "db",
		Username: env.Require("POSTGRES_USER"),
		Password: env.Require("POSTGRES_PASSWORD"),
		DbName:   db.DbName(),
	})
}
