package projects

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	"github.com/troptropcontent/tick_tom/internal/models"
)

type Project struct {
	Name string `form:"project[name]"`
}

func Create(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	var project models.Project
	err := c.Bind(&project)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if project.Name == "" {
		return c.String(http.StatusBadRequest, "name is required")
	}

	project.UserID = current_user.ID
	result := db.DB.Create(&project)
	if result.Error != nil {
		return result.Error
	}

	return c.Redirect(302, c.Echo().Reverse("root"))
}
