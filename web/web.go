package web

import "embed"

//go:embed templates
var TemplatesFS embed.FS

//go:embed all:static/dist
var StaticFS embed.FS

//go:embed locales
var LocalesFS embed.FS
