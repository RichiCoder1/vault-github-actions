package sync

func (ctx *SyncContext) CreateOrUpdateIssue(title string, msg string, id string) error {
	if ctx.repo == "" {
		ctx.logger.Infof("Failed to sync org: %s %s %s", title, msg, id)
	}
	return nil
}
