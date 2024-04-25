package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name   string
	UserID uint
	User   User
}
