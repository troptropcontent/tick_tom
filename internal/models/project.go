package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name   string `form:"project[name]"`
	UserID uint
	User   User
}
