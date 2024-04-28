package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	"github.com/troptropcontent/tick_tom/internal/models"
)

type ShowTemplateData struct {
	HeaderData struct {
		Title string
	}
	Project models.Project
}

func Show(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	id := c.Param("id")
	var project models.Project
	err := db.DB.Where(&models.Project{UserID: current_user.ID}).First(&project, id).Error
	if err != nil {
		return err
	}

	return c.Render(200, "projects/show.html", ShowTemplateData{
		HeaderData: struct{ Title string }{Title: "Project " + project.Name},
		Project:    project,
	})
}
