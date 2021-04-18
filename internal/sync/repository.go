package sync

import (
	"fmt"
	"sync"

	"github.com/google/go-github/github"
)

var repoLocks sync.Map

func init() {
	repoLocks = sync.Map{}
}

func (ctx *SyncContext) SyncRepository() error {
	loaded, _ := repoLocks.LoadOrStore(fmt.Sprintf("%s/%s", ctx.owner, ctx.repo), &sync.RWMutex{})
	lock := loaded.(*sync.RWMutex)

	lock.RLock()
	config, err := ctx.getRepositoryConfig()
	lock.RUnlock()
	if err != nil {
		return err
	}

	if config == nil {
		return nil
	}

	if config.Secrets == nil {
		ctx.logger.Debugf("No secrets found for repository")
		return nil
	}

	lock.Lock()
	defer lock.Unlock()

	return nil
}

func (ctx *SyncContext) getRepositoryConfig() (*RepositoryConfiguration, error) {
	file, dir, res, err := ctx.github.Repositories.GetContents(
		ctx.ctx,
		ctx.owner,
		ctx.repo,
		ctx.config.RepoConfigurationFile,
		// Get from default ref
		&github.RepositoryContentGetOptions{},
	)
	if res != nil && err != nil {
		if res.StatusCode == 404 {
			ctx.logger.Debugf("No config file found for repo: %s", ctx.config.RepoConfigurationFile)
			return nil, nil
		}
		return nil, err
	} else if err != nil {
		return nil, err
	}

	if dir != nil {
		ctx.logger.Errorf("Specified configuration file is a directory: %s", ctx.config.RepoConfigurationFile)
		return nil, err
	}

	fileContents, err := file.GetContent()
	if err != nil {
		return nil, err
	}

	config, err := ParseRepoConfig(fileContents)
	if err != nil {
		return nil, err
	}
	return config, nil
}
