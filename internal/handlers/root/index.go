package root

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
)

func Index(c echo.Context) error {
	current_user_email, _ := auth.GetCurrentUserEmailFromContext(c)
	return c.Render(200, "root/index.html", map[string]interface{}{
		"current_user_email": current_user_email,
	})
}
