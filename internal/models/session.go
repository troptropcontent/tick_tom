package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	HolderID   uint
	HolderType string
	UserID     uint
	StartedAt  time.Time
	EndedAt    time.Time
}

func (s Session) Status() string {
	if s.ID == 0 {
		return "new"
	}
	if s.EndedAt.IsZero() {
		return "in_progress"
	}
	return "stopped"
}

func (s Session) IsNew() bool {
	return s.Status() == "new"
}

func (s Session) IsInProgress() bool {
	return s.Status() == "in_progress"
}

func (s Session) IsStopped() bool {
	return s.Status() == "stopped"
}

func (s Session) TimeSpent() time.Duration {
	if s.IsInProgress() {
		return time.Since(s.StartedAt)
	}
	return s.EndedAt.Sub(s.StartedAt)
}
