package models

import "gorm.io/gorm"

// This method is used to migrate the database
func Init(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Project{},
		&Task{},
		&Session{},
	)
}
