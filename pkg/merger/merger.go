package merger

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/nullify-platform/config-file-parser/pkg/validator"
)

func MergeConfigFiles(
	globalConfig *models.Configuration,
	repoConfig *models.Configuration,
) *models.Configuration {
	config := parser.NewDefaultConfig()

	// override with global config

	if globalConfig != nil {
		if globalConfig.FailBuilds != nil {
			config.FailBuilds = globalConfig.FailBuilds
		}

		if globalConfig.SeverityThreshold != "" && validator.ValidateSeverityThreshold(config) {
			config.SeverityThreshold = globalConfig.SeverityThreshold
		}

		config.IgnoreDirs = globalConfig.IgnoreDirs
		config.IgnorePaths = globalConfig.IgnorePaths

		config.Code.Ignore = globalConfig.Code.Ignore
		config.Dependencies.Ignore = globalConfig.Dependencies.Ignore
		config.Secrets.Ignore = globalConfig.Secrets.Ignore

		if len(globalConfig.Notifications) > 0 && config.Notifications == nil {
			config.Notifications = globalConfig.Notifications
		}

		for k, v := range globalConfig.Notifications {
			config.Notifications[k] = v
		}

		if len(globalConfig.ScheduledNotifications) > 0 && config.ScheduledNotifications == nil {
			config.ScheduledNotifications = globalConfig.ScheduledNotifications
		}

		for k, v := range globalConfig.ScheduledNotifications {
			config.ScheduledNotifications[k] = v
		}
	}

	// override with repo config

	if repoConfig == nil {
		return config
	}

	if repoConfig.FailBuilds != nil {
		config.FailBuilds = repoConfig.FailBuilds
	}

	if repoConfig.SeverityThreshold != "" && validator.ValidateSeverityThreshold(config) {
		config.SeverityThreshold = repoConfig.SeverityThreshold
	}

	if len(repoConfig.IgnoreDirs) > 0 {
		config.IgnoreDirs = repoConfig.IgnoreDirs
	}

	if len(repoConfig.IgnorePaths) > 0 {
		config.IgnorePaths = repoConfig.IgnorePaths
	}

	if len(repoConfig.Code.Ignore) > 0 {
		config.Code.Ignore = repoConfig.Code.Ignore
	}

	if len(repoConfig.Dependencies.Ignore) > 0 {
		config.Dependencies.Ignore = repoConfig.Dependencies.Ignore
	}

	if len(repoConfig.Secrets.Ignore) > 0 {
		config.Secrets.Ignore = repoConfig.Secrets.Ignore
	}

	if len(repoConfig.Notifications) > 0 && config.Notifications == nil {
		config.Notifications = repoConfig.Notifications
	}

	for k, v := range repoConfig.Notifications {
		config.Notifications[k] = v
	}

	if len(repoConfig.ScheduledNotifications) > 0 && config.ScheduledNotifications == nil {
		config.ScheduledNotifications = repoConfig.ScheduledNotifications
	}

	for k, v := range repoConfig.ScheduledNotifications {
		config.ScheduledNotifications[k] = v
	}

	return config
}
