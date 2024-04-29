package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/troptropcontent/tick_tom/internal/env"
	"github.com/troptropcontent/tick_tom/internal/handlers/auth"
	projects_handlers "github.com/troptropcontent/tick_tom/internal/handlers/projects"
	"github.com/troptropcontent/tick_tom/internal/handlers/root"
	db_initializer "github.com/troptropcontent/tick_tom/internal/initializers/db"
	env_initializer "github.com/troptropcontent/tick_tom/internal/initializers/env"
	models_initializer "github.com/troptropcontent/tick_tom/internal/initializers/models"
)

func init() {
	env_initializer.Init()
	db_initializer.Init()
	models_initializer.Init()
}

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "application/layout.html", data)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(env.Require("GOLANG_MASTER_KEY")))))

	// Auth
	store := sessions.NewCookieStore([]byte(env.Require("GOLANG_MASTER_KEY")))
	store.MaxAge(3600)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = false
	gothic.Store = store

	goth.UseProviders(
		google.New(env.Require("GOOGLE_CLIENT_ID"), env.Require("GOOGLE_CLIENT_SECRET"), "http://localhost:3000/auth/callback?provider=google", "email", "profile"),
	)

	// Renderer
	templates := make(map[string]*template.Template)
	templates["root/index.html"] = template.Must(template.ParseFiles("internal/views/root/index.html", "internal/views/application/layout.html", "internal/views/components/navbar.html", "internal/views/components/header.html"))
	templates["auth/login.html"] = template.Must(template.ParseFiles("internal/views/auth/login.html", "internal/views/application/layout.html", "internal/views/components/navbar.html"))
	templates["projects/index.html"] = template.Must(template.ParseFiles("internal/views/projects/index.html", "internal/views/projects/_project.html", "internal/views/application/layout.html", "internal/views/components/navbar.html", "internal/views/components/header.html"))
	templates["projects/new.html"] = template.Must(template.ParseFiles("internal/views/projects/new.html", "internal/views/application/layout.html", "internal/views/components/navbar.html", "internal/views/components/header.html"))
	templates["projects/show.html"] = template.Must(template.ParseFiles("internal/views/projects/show.html", "internal/views/application/layout.html", "internal/views/components/navbar.html", "internal/views/components/header.html"))
	e.Renderer = &Template{
		templates: templates,
	}

	// Routes
	// Static assets
	e.Static("/static", "public")
	// Authentification
	e.GET("/auth/callback", auth.Signin).Name = "auth.callback"
	e.GET("/auth", auth.OAuth).Name = "auth.oauth"
	e.GET("/auth/login", auth.Login).Name = "auth.login"
	e.POST("/auth/logout", auth.Logout).Name = "auth.logout"

	// Root
	e.GET("/", auth.RequireAuthenticatedUser(root.Index)).Name = "root"

	// Projects
	projects := e.Group("projects")
	projects.Use(auth.RequireAuthenticatedUser)
	projects.GET("/new", projects_handlers.New).Name = "projects.new"
	projects.POST("", projects_handlers.Create).Name = "projects.new"
	projects.GET("", projects_handlers.Index).Name = "projects.index"
	projects.GET("/:id", projects_handlers.Show).Name = "projetcs.show"

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
