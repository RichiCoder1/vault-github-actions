package sync

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/richicoder1/vault-github-actions/internal"
	log "github.com/sirupsen/logrus"
)

type SyncContext struct {
	ctx    context.Context
	github *github.Client
	logger *log.Entry
	config internal.VgaConfig

	owner string
	repo  string
}
