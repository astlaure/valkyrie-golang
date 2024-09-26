package core

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/astlaure/valkyrie-golang/web"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var DefaultLanguage = "en"
var SupportedLanguages = []string{"en", "fr"}
var Bundle = i18n.NewBundle(language.English)

func SetupI18n() {
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	Bundle.LoadMessageFileFS(web.LocalesFS, "locales/en.json")
	Bundle.LoadMessageFileFS(web.LocalesFS, "locales/fr.json")
}

func I18nMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		locale := DefaultLanguage
		sections := strings.Split(c.Request().RequestURI, "/")

		if len(sections) >= 2 && slices.Contains(SupportedLanguages, sections[1]) {
			locale = sections[1]
		}

		c.Set("Locale", locale)
		c.Set("Localizer", i18n.NewLocalizer(Bundle, locale))
		return next(c)
	}
}

// func Trans(c echo.Context) func {
//     func (key string, data map[string]interface{}) string {
//         localizer := c.Get("Localizer").(*i18n.Localizer)
//         message, _ := localizer.Localize(&i18n.LocalizeConfig{
//             MessageID: key,
//         })
//         return message
//     }
// }
