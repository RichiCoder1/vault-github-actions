package sync

import (
	vault "github.com/hashicorp/vault/api"
)

func (ctx *SyncContext) GetVaultClient() (*vault.Client, error) {
	config := vault.DefaultConfig()
	return vault.NewClient(config)
}
