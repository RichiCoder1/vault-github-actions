package web

import (
	"net/url"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/richicoder1/vault-github-actions/internal"
)

func RegisterDev(e *echo.Echo) internal.OnShutdown {
	ui := e.Group("/ui")

	devServer := exec.Command("npm", "run", "dev")
	devServer.Dir = "web"
	devServer.Stdout = os.Stdout
	devServer.Stderr = os.Stderr

	err := devServer.Start()
	if err != nil {
		e.Logger.Fatal("Failed to start dev server", err)
		return func() error { return nil }
	}

	devServerUri, err := url.Parse("http://localhost:3000")
	if err != nil {
		e.Logger.Fatal("Failed to start dev server: ", err)
		return func() error { return nil }
	}

	targets := []*middleware.ProxyTarget{
		{
			URL: devServerUri,
		},
	}

	ui.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	return func() error {
		return devServer.Process.Kill()
	}
}
