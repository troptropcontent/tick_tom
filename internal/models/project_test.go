package models

import (
	"testing"
	"time"

	"github.com/troptropcontent/tick_tom/db"
	db_initializer "github.com/troptropcontent/tick_tom/internal/initializers/db"
	env_initializer "github.com/troptropcontent/tick_tom/internal/initializers/env"
)

func TestProject_TotalTimeSpent(t *testing.T) {
	env_initializer.Init()
	db_initializer.Init()
	Init(db.DB)
	t.Cleanup(func() { db.EmptyTables() })

	user := User{
		Email: "test@test.com",
	}
	db.DB.Create(&user)

	project := Project{
		UserID: user.ID,
		Name:   "test",
	}
	db.DB.Create(&project)

	type fields struct {
		Project  Project
		Sessions []Session
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "When there is no session",
			fields: fields{
				Project: project,
			},
			want: 0,
		},
		{
			name: "When there is some sessions",
			fields: fields{
				Project: project,
				Sessions: []Session{
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
						EndedAt:    time.Now().Add(1 * time.Hour),
					},
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
						EndedAt:    time.Now().Add(1 * time.Hour),
					},
				},
			},
			want: 2 * time.Hour,
		},
		{
			name: "When there is an session that is not ended",
			fields: fields{
				Project: project,
				Sessions: []Session{
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
						EndedAt:    time.Now().Add(1 * time.Hour),
					},
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
						EndedAt:    time.Now().Add(1 * time.Hour),
					},
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
					},
				},
			},
			want: 2 * time.Hour,
		},
		{
			name: "When there is sessions of all types of duration",
			fields: fields{
				Project: project,
				Sessions: []Session{
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
						EndedAt:    time.Now().Add(1 * time.Hour),
					},
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
						EndedAt:    time.Now().Add(1 * time.Minute),
					},
					{
						HolderID:   project.ID,
						HolderType: "projects",
						StartedAt:  time.Now(),
						EndedAt:    time.Now().Add(1 * time.Second),
					},
				},
			},
			want: 1*time.Hour + 1*time.Minute + 1*time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.fields.Sessions) > 0 {
				db.DB.Create(&tt.fields.Sessions)
			}
			if got := tt.fields.Project.TotalTimeSpent(); got != tt.want {
				t.Errorf("Project.TotalTimeSpent() = %v, want %v", got, tt.want)
			}
			db.EmptyTables("sessions")
		})
	}
}
