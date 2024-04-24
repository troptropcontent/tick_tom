package auth

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func Signin(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	sess, _ := session.Get(SESSION_ID, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	}
	sess.Values[SESSION_KEY_EMAIL] = user.Email
	sess.Save(c.Request(), c.Response())

	if err != nil {
		fmt.Println("Error:", err)
		return c.Render(500, "auth/error.html", nil)
	}

	return c.Render(200, "auth/success.html", user)
}
