package db_initializer

import (
	"github.com/troptropcontent/tick_tom/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	dsn := "host=db user=postgres password=postgres dbname=postgres"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.DB = database
}
