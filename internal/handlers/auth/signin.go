package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/troptropcontent/tick_tom/db"
	"github.com/troptropcontent/tick_tom/internal/models"
)

func findOrCreateUser(email string) (*models.User, error) {
	var user models.User
	result := db.DB.FirstOrCreate(&user, models.User{Email: email})
	return &user, result.Error
}

func saveUserIdInSession(c echo.Context, user *models.User) error {
	sess, err := session.Get(SESSION_ID, c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	}
	fmt.Println("setting user ID:", user.ID)
	sess.Values[SESSION_KEY_USER_ID] = user.ID
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return nil
}

func Signin(c echo.Context) error {
	authenticated_user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Render(500, "auth/error.html", nil)
	}

	fmt.Println("Authenticated user:", authenticated_user.Email)
	user, err := findOrCreateUser(authenticated_user.Email)
	if err != nil {
		return c.Render(500, "auth/error.html", nil)
	}
	fmt.Printf("Persisted user: %#v\n", user)
	if err := saveUserIdInSession(c, user); err != nil {
		return c.Render(500, "auth/error.html", nil)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
