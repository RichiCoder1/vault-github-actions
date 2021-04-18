package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"
	"github.com/cristalhq/aconfig/aconfigyaml"

	"github.com/richicoder1/vault-github-actions/internal"
	"github.com/richicoder1/vault-github-actions/internal/logging"
	"github.com/richicoder1/vault-github-actions/internal/web"
	"github.com/richicoder1/vault-github-actions/internal/webhook"

	bus "github.com/mustafaturan/bus/v2"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"

	log "github.com/sirupsen/logrus"
)

var DEBUG = "1"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	var cfg internal.VgaConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		EnvPrefix:  "VGA",
		FlagPrefix: "vga",
		FileFlag:   "config",
		MergeFiles: true,
		Files: []string{
			"/var/opt/vga/config.yaml",
			"/var/opt/vga/config.toml",
			"vga.config.yaml",
			"vga.config.toml",
		},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
			".toml": aconfigtoml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Logger = &logging.LogrusEchoLogger{}

	if cfg.LogLevel != "" {
		lvl, err := log.ParseLevel(cfg.LogLevel)
		if err != nil {
			panic(fmt.Errorf("Failed to parse log level: %w", err))
		}
		log.SetLevel(lvl)
	}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	node := uint64(1)
	initialTime := uint64(1577865600000) // set 2020-01-01 PST as initial time
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		panic(err)
	}

	// init an id generator
	var idGenerator bus.Next = (*m).Next

	// create a new bus instance
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}

	webhookConfig, err := webhook.GetConfig(&cfg)
	if err != nil {
		if errors.Is(err, webhook.MissingConfigError) {
			e.GET("/", func(c echo.Context) error {
				return c.Redirect(http.StatusTemporaryRedirect, "/ui/setup")
			})
		}
	} else {
		webhook.Use(e, b, webhookConfig, DEBUG == "1")
	}

	shutdownCbs := []internal.OnShutdown{}
	if DEBUG == "1" {
		shutdownCbs = append(shutdownCbs, web.RegisterDev(e))
	} else {
		web.RegisterStatic(e)
	}

	go func() {
		if err := e.Start(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server:\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	e.Logger.Info("Got interrupt signal")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, cb := range shutdownCbs {
		if err := cb(); err != nil {
			e.Logger.Fatal(err)
		}
	}

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
