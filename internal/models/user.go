package models

import (
	"github.com/troptropcontent/tick_tom/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string
}

func (u *User) Create() error {
	db.DB.Create(u)
	return nil
}
