package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	"github.com/troptropcontent/tick_tom/internal/models"
)

type ShowTemplateTranslation struct {
	ButtonMain    string
	ButtonCaption string
	Hours         string
	Minutes       string
	Secondes      string
}
type ShowTemplateData struct {
	HeaderData struct {
		Title string
	}
	Project     models.Project
	LastSession models.Session
	Translation ShowTemplateTranslation
}

func Show(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	id := c.Param("id")
	var project models.Project
	err := db.DB.Where(&models.Project{UserID: current_user.ID}).First(&project, id).Error
	if err != nil {
		return err
	}
	var last_session models.Session
	err = db.DB.Where(&models.Session{
		HolderID: project.ID,
	}).Limit(1).Find(&last_session).Error
	if err != nil {
		return err
	}

	return c.Render(200, "projects/show.html", ShowTemplateData{
		HeaderData:  struct{ Title string }{Title: "Project " + project.Name},
		Project:     project,
		LastSession: last_session,
		Translation: ShowTemplateTranslation{
			ButtonMain: "Start",
		},
	})
}
