package root

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
)

func Index(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)
	return c.Render(200, "root/index.html", current_user)
}
