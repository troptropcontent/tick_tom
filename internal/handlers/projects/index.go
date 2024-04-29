package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	"github.com/troptropcontent/tick_tom/internal/models"
)

type IndexTemplateData struct {
	HeaderData struct {
		Title string
	}
	Projects []struct {
		Instance models.Project
		ShowUrl  string
	}
}

func buildIndexTemplateData(c echo.Context, projects []models.Project) IndexTemplateData {
	template_data := IndexTemplateData{
		HeaderData: struct {
			Title string
		}{
			Title: "All your projects",
		},
	}

	for _, project := range projects {
		template_data.Projects = append(template_data.Projects, struct {
			Instance models.Project
			ShowUrl  string
		}{
			Instance: project,
			ShowUrl:  c.Echo().Reverse("projetcs.show", project.ID),
		},
		)
	}
	return template_data
}

func Index(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	var projects []models.Project
	err := db.DB.Where(&models.Project{
		UserID: current_user.ID,
	}).Find(&projects).Error
	if err != nil {
		return err
	}

	return c.Render(
		200,
		"projects/index.html",
		buildIndexTemplateData(c, projects),
	)
}
