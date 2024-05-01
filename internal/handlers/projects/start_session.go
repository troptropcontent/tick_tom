package projects

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	"github.com/troptropcontent/tick_tom/internal/models"
)

func StartSession(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	project_id, _ := strconv.Atoi(c.Param("id"))
	project := models.Project{}
	db.DB.Where(&models.Project{UserID: current_user.ID}).First(&project, project_id)
	session := models.Session{}
	project.LastSession(current_user.ID, &session)

	if session.IsInProgress() {
		template_data, err := loadShowTemplateData(project)
		if err != nil {
			return err
		}

		template_data.Name = project.Name
		template_data.HeaderData.Title = project.Name
		template_data.ProjectID = project.ID
		template_data.Error = "Session is already in progress"
		return c.Render(http.StatusBadRequest, "projects/show.html", template_data)
	}

	new_session := models.Session{
		HolderID:   project.ID,
		HolderType: "projects",
		UserID:     current_user.ID,
		StartedAt:  time.Now().UTC(),
	}
	db.DB.Create(&new_session)

	template_data, err := loadShowTemplateData(project)
	if err != nil {
		return err
	}

	template_data.Name = project.Name
	template_data.HeaderData.Title = project.Name
	template_data.ProjectID = project.ID

	return c.Render(200, "projects/start_stop_session.partial.html", template_data)
}
