package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	duration_helpers "github.com/troptropcontent/tick_tom/internal/helpers/duration"
	"github.com/troptropcontent/tick_tom/internal/models"
)

type ShowTemplateTranslation struct {
	ButtonMain string
}
type ShowTemplateData struct {
	HeaderData struct {
		Title string
	}
	Name         string
	Duration     string
	IsInProgress bool
	Translation  ShowTemplateTranslation
}

func loadProject(user_id uint, id string) (models.Project, error) {
	var project models.Project
	err := db.DB.Where(&models.Project{UserID: user_id}).First(&project, id).Error
	return project, err
}

func loadTemplateData(project models.Project) (ShowTemplateData, error) {
	var last_session models.Session
	err := project.LastSession(project.UserID, &last_session)
	if err != nil {
		return ShowTemplateData{}, err
	}

	if last_session.IsInProgress() {
		return ShowTemplateData{
			IsInProgress: true,
			Duration:     duration_helpers.RjustDuration(last_session.TimeSpent()),
			Translation: ShowTemplateTranslation{
				ButtonMain: "Stop",
			},
		}, nil
	}

	return ShowTemplateData{
		IsInProgress: false,
		Duration:     duration_helpers.RjustDuration(project.TotalTimeSpent()),
		Translation: ShowTemplateTranslation{
			ButtonMain: "Start",
		},
	}, nil
}

func Show(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	id := c.Param("id")
	project, err := loadProject(current_user.ID, id)
	if err != nil {
		return err
	}

	template_data, err := loadTemplateData(project)
	if err != nil {
		return err
	}

	template_data.Name = project.Name
	template_data.HeaderData.Title = project.Name

	return c.Render(200, "projects/show.html", template_data)
}
