package webhook

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/google/go-github/github"

	bus "github.com/mustafaturan/bus/v2"
)

func Use(e *echo.Echo, b *bus.Bus, config VgaWebhookConfig, debug bool) error {
	e.POST("/webhook", func(c echo.Context) error {
		r := c.Request()

		payload, err := github.ValidatePayload(r, []byte(config.WebhookSecret))
		if err != nil {
			return err
		}

		event, err := github.ParseWebHook(github.WebHookType(r), payload)
		if err != nil {
			return err
		}

		eventType := event.(github.Event)
		b.RegisterTopics(*eventType.Type)

		_, err = b.Emit(context.Background(), *eventType.Type, event)
		if err != nil {
			return err
		}

		return c.String(http.StatusAccepted, "OK")
	})
	return nil
}
