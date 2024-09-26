package core

import (
	"errors"
	"html/template"
	"io"
	"io/fs"
	"strings"

	"github.com/astlaure/valkyrie-golang/web"
	"github.com/labstack/echo/v4"
)

type (
	RendererConfig struct {
		Root          string
		Extension     string
		LayoutFolder  string
		DefaultLayout string
	}

	Template struct {
		templates map[string]*template.Template
		config    RendererConfig
	}
)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		return errors.New("Template not found -> " + name)
	}

	layout, ok := c.Get("layout").(string)

	if !ok {
		layout = t.config.DefaultLayout
	}

	return tmpl.ExecuteTemplate(w, layout, data)
}

func registerTemplates(config RendererConfig) map[string]*template.Template {
	var elements = make(map[string]*template.Template)
	var layouts []string

	// Load the layouts
	fs.WalkDir(web.TemplatesFS, config.Root+"/"+config.LayoutFolder, func(path string, d fs.DirEntry, err error) error {
		path = strings.ReplaceAll(path, "\\", "/") // Windows path fix

		if !d.IsDir() && strings.HasSuffix(path, config.Extension) {
			layouts = append(layouts, path)
		}

		return nil
	})

	// Load the templates
	fs.WalkDir(web.TemplatesFS, config.Root, func(path string, d fs.DirEntry, err error) error {
		path = strings.ReplaceAll(path, "\\", "/") // Windows path fix

		if !d.IsDir() && strings.HasSuffix(path, config.Extension) {
			var name, _ = strings.CutPrefix(
				strings.ReplaceAll(path, config.Root, ""),
				"/",
			)

			var files = append(layouts, path)
			elements[name] = template.Must(template.ParseFS(web.TemplatesFS, files...))
		}

		return nil
	})

	return elements
}

// func Renderer(config RendererConfig) *Template {
// 	return &Template{
// 		templates: registerTemplates(config),
// 		config:    config,
// 	}
// }

var config = RendererConfig{
	Root:          "templates",
	Extension:     ".html",
	LayoutFolder:  "layouts",
	DefaultLayout: "main",
}

var Renderer = &Template{
	templates: registerTemplates(config),
	config:    config,
}

// var Renderer = app.Renderer = core.Renderer(core.RendererConfig{
//     Root:          "views",
//     Extension:     ".html",
//     LayoutFolder:  "layouts",
//     DefaultLayout: "main",
// })
