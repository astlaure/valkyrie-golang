package users

import (
	"net/http"

	"github.com/astlaure/valkyrie-golang/internal/core"
	"github.com/labstack/echo/v4"
)

func handleGetUsers(c echo.Context) error {
	model := core.TemplateModel(c)
	return c.Render(http.StatusOK, "users/index.html", model)
}

func RegisterGroup(prefix string, app *echo.Echo) {
	group := app.Group(prefix)

	group.GET("", handleGetUsers)
}
