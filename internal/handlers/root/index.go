package root

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
)

func Index(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	logout_url := c.Echo().Reverse("auth.logout")
	return c.Render(200, "root/index.html", map[string]any{
		"current_user": current_user,
		"logout_url":   logout_url,
	})
}
