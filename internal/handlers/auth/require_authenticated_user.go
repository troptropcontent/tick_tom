package auth

import (
	"errors"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetCurrentUserEmailFromCookie(c echo.Context) (string, error) {
	sess, err := session.Get(SESSION_ID, c)
	if err != nil {
		return "", err
	}
	current_user_email := sess.Values[SESSION_KEY_EMAIL]
	if current_user_email == nil {
		return "", errors.New("user is not authenticated")
	}
	return current_user_email.(string), nil
}

func GetCurrentUserEmailFromContext(c echo.Context) (string, error) {
	data := c.Get("current_user_email")
	if data == nil {
		return "", errors.New("user is not authenticated")
	}
	return data.(string), nil
}

func SetCurrentUserEmailInContext(c echo.Context, email string) {
	c.Set("current_user_email", email)
}

func RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		current_user_email_from_cookie, _ := GetCurrentUserEmailFromCookie(c)
		if current_user_email_from_cookie == "" {
			return c.Redirect(302, "/auth/login")
		}
		SetCurrentUserEmailInContext(c, current_user_email_from_cookie)

		return next(c)
	}
}
