package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Task      Task
	TaskID    uint
	UserID    uint
	StartedAt time.Time
	EndedAt   time.Time
}
