package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func OAuth(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}
