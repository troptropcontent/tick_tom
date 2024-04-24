package models_initializer

import (
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/models"
)

func Init() {
	db.DB.AutoMigrate(&models.User{})
}
