package webhook

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/richicoder1/vault-github-actions/internal"
)

type VgaWebhookConfig struct {
	WebhookSecret  string
	PrivateKey     []byte
	AppId          int64
	InstallationId int64
}

var MissingConfigError = errors.New("missing config")

func GetConfig(config *internal.VgaConfig) (VgaWebhookConfig, error) {
	if config.WebhookSecret == "" && config.WebhookSecretFile == "" {
		return VgaWebhookConfig{}, fmt.Errorf("No webhook secret provided: %w", MissingConfigError)
	}

	webhookSecret := config.WebhookSecret
	if config.WebhookSecretFile != "" {
		data, err := ioutil.ReadFile(config.WebhookSecretFile)
		if err != nil {
			return VgaWebhookConfig{}, fmt.Errorf("Unable to read webhook secret file: %w", err)
		}

		webhookSecret = strings.TrimSpace(string(data))
	}

	if config.Github.PrivateKey == "" && config.Github.PrivateKeyFile == "" {
		return VgaWebhookConfig{}, fmt.Errorf("No private key provided: %w", MissingConfigError)
	}

	var privateKeyPem []byte
	if config.Github.PrivateKeyFile != "" {
		data, err := ioutil.ReadFile(config.Github.PrivateKeyFile)
		if err != nil {
			return VgaWebhookConfig{}, fmt.Errorf("Unable to read private key file: %w", err)
		}

		privateKeyPem = data
	} else {
		privateKeyDecoded, err := base64.StdEncoding.DecodeString(config.Github.PrivateKey)
		if err != nil {
			return VgaWebhookConfig{}, fmt.Errorf("Unable to decode private key: %w", err)
		}
		privateKeyPem = privateKeyDecoded
	}

	if config.Github.AppId == 0 {
		return VgaWebhookConfig{}, fmt.Errorf("Missing App ID: %s", MissingConfigError)
	}

	return VgaWebhookConfig{
		WebhookSecret: webhookSecret,
		PrivateKey:    privateKeyPem,
		AppId:         config.Github.AppId,
	}, nil
}
