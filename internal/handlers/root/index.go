package root

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	"github.com/troptropcontent/tick_tom/internal/models"
)

func Index(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	var projects []models.Project
	db.DB.Where(&models.Project{UserID: current_user.ID}).Find(&projects)
	if len(projects) == 0 {
		c.Redirect(302, c.Echo().Reverse("projects.new"))
	}

	return c.Render(200, "root/index.html", map[string]string{
		"name": "Index",
	})
}
