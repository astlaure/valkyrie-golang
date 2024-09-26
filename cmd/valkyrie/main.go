package main

import (
	"github.com/astlaure/valkyrie-golang/internal/app"
	"github.com/astlaure/valkyrie-golang/internal/core"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	core.SetupDatabase()
	core.SetupI18n()
	app.Bootstrap()
}
