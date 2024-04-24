package auth

import "github.com/labstack/echo/v4"

func Login(c echo.Context) error {
	return c.Render(200, "auth/login.html", nil)
}
