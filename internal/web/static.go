package web

import (
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static
var staticFiles embed.FS

func RegisterStatic(e *echo.Echo) {
	ui := e.Group("/ui")

	ui.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "static",
		HTML5: true,
	}))
}
