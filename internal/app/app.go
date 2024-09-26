package app

import (
	"log"
	"net/http"

	"github.com/astlaure/valkyrie-golang/internal/core"
	"github.com/astlaure/valkyrie-golang/internal/users"
	"github.com/astlaure/valkyrie-golang/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Bootstrap() {
	app := echo.New()

	app.Debug = core.Debug
	app.Renderer = core.Renderer
	app.Validator = core.Validator

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(web.StaticFS),
		Root:       "static/dist",
	}))
	app.Use(core.I18nMiddleware)

	app.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/"+core.DefaultLanguage)
	})

	users.RegisterGroup(core.GetPrefix("/users"), app)

	app.GET(core.GetPrefix(""), func(c echo.Context) error {
		model := core.TemplateModel(c)
		return c.Render(http.StatusOK, "index.html", model)
	})

	log.Fatal(app.Start("127.0.0.1:8080"))
}
