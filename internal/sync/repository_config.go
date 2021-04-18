package sync

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type RepositoryConfiguration struct {
	Secrets *map[string]SecretRequest `yaml:"secrets,omitempty"`
}

type SecretRequest struct {
	Path     string `yaml:"path"`
	Selector string `yaml:"selector"`
	MaxAge   string `yaml:"maxAge,omitempty"`
}

func ParseRepoConfig(configFile string) (*RepositoryConfiguration, error) {
	config := RepositoryConfiguration{}
	if err := yaml.Unmarshal([]byte(configFile), &config); err != nil {
		return nil, fmt.Errorf("Failed to parse configuration file: %w", err)
	}

	return &config, nil
}
