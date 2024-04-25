package models

import (
	"github.com/troptropcontent/tick_tom/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
}

func (u *User) Create() error {
	return db.DB.Create(u).Error
}
