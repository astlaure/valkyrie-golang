package core

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var Debug = os.Getenv("DEBUG") == "true"

var Env = os.Getenv("ENV")

var IsProduction = os.Getenv("ENV") == "production"

var BasePath = "/:locale"

func TemplateModel(c echo.Context) echo.Map {
	return echo.Map{
		"Env": Env,
		"T": func(key string, data map[string]interface{}) string {
			localizer := c.Get("Localizer").(*i18n.Localizer)
			message, _ := localizer.Localize(&i18n.LocalizeConfig{
				MessageID: key,
			})
			return message
		},
		"Assets": GetViteAssets(),
	}
}

func GetPrefix(prefix string) string {
	if prefix == "/" || prefix == "" {
		return BasePath
	}

	if !strings.HasPrefix(prefix, "/") {
		return BasePath + "/" + prefix
	}

	return BasePath + prefix
}
