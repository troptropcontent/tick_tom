package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
)

func New(c echo.Context) error {
	current_user, _ := auth.GetCurrentUserFromContext(c)

	return c.Render(200, "projects/new.html", map[string]any{
		"current_user": current_user,
		"headerData": map[string]any{
			"title": "Create a new project",
		},
		"translations": map[string]any{
			"form": map[string]any{
				"inputs": map[string]any{
					"name": map[string]any{
						"label":       "Name of the project :",
						"placeholder": "Name of the project",
					},
					"submit": map[string]any{
						"text": "Create",
					},
				},
			},
		},
	})
}
