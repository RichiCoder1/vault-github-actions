package sync

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"

	bus "github.com/mustafaturan/bus/v2"

	"github.com/richicoder1/vault-github-actions/internal"
	"github.com/richicoder1/vault-github-actions/internal/webhook"
)

func RegisterHandlers(b *bus.Bus, logger *log.Entry, appConfig internal.VgaConfig, config webhook.VgaWebhookConfig) {
	repositoryHandler := bus.Handler{
		Handle: func(ctx context.Context, e *bus.Event) {
			pushEvent := e.Data.(github.PushEvent)

			githubClient, err := GetInstallationClient(&config, *pushEvent.Installation.ID)
			if err != nil {
				logger.Errorf("Failed to create github client: %w", err)
				return
			}

			owner := *pushEvent.Repo.Owner.Name
			repo := *pushEvent.Repo.Name

			syncLogger := logger.WithFields(log.Fields{
				"repo":  repo,
				"owner": owner,
				"type":  "push",
			})

			syncContext := SyncContext{
				ctx:    ctx,
				github: githubClient,
				logger: syncLogger,
				config: appConfig,
				owner:  owner,
				repo:   repo,
			}

			go func() {
				if err := syncContext.SyncRepository(); err != nil {
					syncLogger.WithError(err).Error("Failed to sync repository")
				}
			}()
		},
		Matcher: "^push$",
	}
	b.RegisterHandler("repository_sync", &repositoryHandler)
}

func GetAppClient(config *webhook.VgaWebhookConfig) (*github.Client, error) {
	tr := http.DefaultTransport

	itr, err := ghinstallation.NewAppsTransport(tr, config.AppId, config.PrivateKey)
	if err != nil {
		return nil, err
	}

	return github.NewClient(&http.Client{Transport: itr}), nil
}

func GetInstallationClient(config *webhook.VgaWebhookConfig, installationId int64) (*github.Client, error) {
	tr := http.DefaultTransport

	itr, err := ghinstallation.New(tr, config.AppId, installationId, config.PrivateKey)
	if err != nil {
		return nil, err
	}

	return github.NewClient(&http.Client{Transport: itr}), nil
}
