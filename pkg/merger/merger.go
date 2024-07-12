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

		config.Code.AutoFix = globalConfig.Code.AutoFix
		config.Code.Ignore = globalConfig.Code.Ignore
		config.Dependencies.AutoFix = globalConfig.Dependencies.AutoFix
		config.Dependencies.Ignore = globalConfig.Dependencies.Ignore
		config.Secrets.Ignore = globalConfig.Secrets.Ignore
		config.Secrets.CustomPatterns = globalConfig.Secrets.CustomPatterns
		config.SecretsWhitelist = globalConfig.SecretsWhitelist

		if globalConfig.Integrations.Jira != nil {
			config.Integrations.Jira = globalConfig.Integrations.Jira
		}

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

	if repoConfig.Integrations.Jira != nil {
		config.Integrations.Jira = repoConfig.Integrations.Jira
		if repoConfig.Integrations.Jira.Priorities != nil {
			config.Integrations.Jira.Priorities = repoConfig.Integrations.Jira.Priorities
		} else if globalConfig != nil && globalConfig.Integrations.Jira != nil && globalConfig.Integrations.Jira.Priorities != nil {
			// copying over repoConfig if globalConfig is set may overwrite the priorities,
			// so we need to copy over the globalConfig priorities again if repo level priorities are not set
			config.Integrations.Jira.Priorities = globalConfig.Integrations.Jira.Priorities
		}

		if repoConfig.Integrations.Jira.Assignee != nil {
			config.Integrations.Jira.Assignee = repoConfig.Integrations.Jira.Assignee
		} else if globalConfig != nil && globalConfig.Integrations.Jira != nil && globalConfig.Integrations.Jira.Assignee != nil {
			// assignee is set in the global config, but not in the repo config
			config.Integrations.Jira.Assignee = globalConfig.Integrations.Jira.Assignee
		}
	}

	if len(repoConfig.IgnoreDirs) > 0 {
		config.IgnoreDirs = repoConfig.IgnoreDirs
	}

	if len(repoConfig.IgnorePaths) > 0 {
		config.IgnorePaths = repoConfig.IgnorePaths
	}

	if repoConfig.Code.AutoFix != nil {
		config.Code.AutoFix = repoConfig.Code.AutoFix
	}

	if len(repoConfig.Code.Ignore) > 0 {
		config.Code.Ignore = repoConfig.Code.Ignore
	}

	if repoConfig.Dependencies.AutoFix != nil {
		config.Dependencies.AutoFix = repoConfig.Dependencies.AutoFix
	}

	if len(repoConfig.Dependencies.Ignore) > 0 {
		config.Dependencies.Ignore = repoConfig.Dependencies.Ignore
	}

	if len(repoConfig.Secrets.Ignore) > 0 {
		config.Secrets.Ignore = repoConfig.Secrets.Ignore
	}

	if len(repoConfig.SecretsWhitelist) > 0 {
		config.SecretsWhitelist = repoConfig.SecretsWhitelist
	}

	config.Secrets.CustomPatternsOverrideGlobal = repoConfig.Secrets.CustomPatternsOverrideGlobal

	if repoConfig.Secrets.CustomPatternsOverrideGlobal {
		// override global custom patterns with repo custom patterns
		config.Secrets.CustomPatterns = repoConfig.Secrets.CustomPatterns
	} else if config.Secrets.CustomPatterns == nil {
		config.Secrets.CustomPatterns = repoConfig.Secrets.CustomPatterns
	} else {
		// merge repo custom patterns with global custom patterns
		for k, v := range repoConfig.Secrets.CustomPatterns {
			config.Secrets.CustomPatterns[k] = v
		}
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
