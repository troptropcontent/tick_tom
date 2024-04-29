package db

import (
	"fmt"

	"github.com/troptropcontent/tick_tom/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DbNamePrefix = "tick_tom"
)

type Config struct {
	Username string
	Password string
	DbName   string
	Host     string
}

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s", c.Host, c.Username, c.Password, c.DbName)
}

func New(config Config) *gorm.DB {
	dsn := config.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var (
	DB *gorm.DB
)

func DbName() string {
	return fmt.Sprintf("%s_%s", DbNamePrefix, env.Current())
}

func Create() {
	DB = New(Config{
		Username: env.Require("POSTGRES_USER"),
		Password: env.Require("POSTGRES_PASSWORD"),
		DbName:   "postgres",
		Host:     "db",
	})

	fmt.Println("Creating database : ", DbName())
	result := DB.Exec(fmt.Sprintf("CREATE DATABASE %s", DbName()))
	if result.Error != nil {
		panic(result.Error)
	}
}
