package auth

import (
	"errors"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/models"
)

func GetCurrentUserIdFromCookie(c echo.Context) (uint, error) {
	sess, err := session.Get(SESSION_ID, c)
	if err != nil {
		return 0, err
	}
	current_user_id := sess.Values[SESSION_KEY_USER_ID]
	if current_user_id == nil {
		return 0, errors.New("user is not authenticated")
	}
	return current_user_id.(uint), nil
}

func GetCurrentUserFromContext(c echo.Context) (models.User, error) {
	data := c.Get(CONTEXT_KEY_CURRENT_USER)
	if data == nil {
		return models.User{}, errors.New("user is not authenticated")
	}
	return data.(models.User), nil
}

func SetCurrentUserInContext(c echo.Context, user models.User) {
	c.Set(CONTEXT_KEY_CURRENT_USER, user)
}

func RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		current_user_id_from_cookie, _ := GetCurrentUserIdFromCookie(c)
		if current_user_id_from_cookie == 0 {
			return c.Redirect(302, "/auth/login")
		}

		var user models.User
		db.DB.Limit(1).Find(&user, current_user_id_from_cookie)
		if user.ID == 0 {
			return c.Redirect(302, "/auth/login")
		}
		SetCurrentUserInContext(c, user)

		return next(c)
	}
}
