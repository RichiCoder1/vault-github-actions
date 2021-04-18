package internal

type VgaConfig struct {
	Address           string `default:"" usage:"The address to listen on"`
	Port              int    `default:"1323" usage:"The port to listen on"`
	WebhookSecret     string `usage:"The Github App webhook secret to verify"`
	WebhookSecretFile string `usage:"The file containing the Github App webhook secret to verify"`
	Github            struct {
		AppId          int64  `usage:"The id of the app"`
		PrivateKey     string `usage:"The base64 encoded private key for the GitHub App"`
		PrivateKeyFile string `usage:"The path to the private key for the GitHub App"`
		InstallationId int64  `usage:"The installation to process. By default runs for all installations"`
	}
	LogLevel string `default:"info" usage:"The log level"`

	Organizations []OrganizationConfig `usage:"The organizations to include in syncing"`

	RepoConfigurationFile string `default:".github/vault.yml" usage:"The configuration file to use for repository requests"`
}

type OrganizationConfig struct {
	Name    string    `usage:"The name of the organization"`
	Include *[]string `usage:"Which repositories to include in syncing"`
	Exclude *[]string `usage:"Which repositories to exclude from syncing"`

	Access *[]SecretAccessRestriction `usage:"A list of secret matchers and repository restrictions"`

	IncludePublic bool `default:"false" usage:"Whether or not to include public repositories in syncs"`
	ShowPrPreview bool `default:"true" usage:"Whether or not to comment on PRs with the result of a change to config file"`
}

type SecretAccessRestriction struct {
	Prefix string    `usage:"The path prefix of a vault request"`
	Regex  string    `usage:"A regex to match against vault requests"`
	Allow  *[]string `usage:"Which repositories to allow to request this secret"`
	Forbid *[]string `usage:"Which repositories to from from requesting this secret"`
}
