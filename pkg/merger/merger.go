package merger

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/validator"
)

func MergeConfigFiles(
	defaultConfg *models.Configuration,
	extraConfigs ...*models.Configuration,
) *models.Configuration {
	if defaultConfg == nil {
		return nil
	}

	// copy default config
	config := *defaultConfg

	// loop over extraConfigs and override settings
	for _, extraConfig := range extraConfigs {
		if extraConfig == nil {
			continue
		}

		if extraConfig.FailBuilds != nil {
			config.FailBuilds = extraConfig.FailBuilds
		}

		if extraConfig.SeverityThreshold != "" && validator.ValidateSeverityThreshold(extraConfig) {
			config.SeverityThreshold = extraConfig.SeverityThreshold
		}

		if extraConfig.Integrations.Jira != nil {
			if config.Integrations.Jira == nil {
				config.Integrations.Jira = extraConfig.Integrations.Jira
			} else {
				if extraConfig.Integrations.Jira.Priorities != nil {
					config.Integrations.Jira.Priorities = extraConfig.Integrations.Jira.Priorities
				}

				if extraConfig.Integrations.Jira.Assignee != nil {
					config.Integrations.Jira.Assignee = extraConfig.Integrations.Jira.Assignee
				}
			}
		}

		if len(extraConfig.IgnoreDirs) > 0 {
			config.IgnoreDirs = extraConfig.IgnoreDirs
		}

		if len(extraConfig.IgnorePaths) > 0 {
			config.IgnorePaths = extraConfig.IgnorePaths
		}

		if extraConfig.Code.AutoFix != nil {
			config.Code.AutoFix = extraConfig.Code.AutoFix
		}

		if len(extraConfig.Code.Ignore) > 0 {
			config.Code.Ignore = extraConfig.Code.Ignore
		}

		if extraConfig.Dependencies.AutoFix != nil {
			config.Dependencies.AutoFix = extraConfig.Dependencies.AutoFix
		}

		if len(extraConfig.Dependencies.Ignore) > 0 {
			config.Dependencies.Ignore = extraConfig.Dependencies.Ignore
		}

		if len(extraConfig.Secrets.Ignore) > 0 {
			config.Secrets.Ignore = extraConfig.Secrets.Ignore
		}

		if len(extraConfig.SecretsWhitelist) > 0 {
			config.SecretsWhitelist = extraConfig.SecretsWhitelist
		}

		config.Secrets.CustomPatternsOverrideGlobal = extraConfig.Secrets.CustomPatternsOverrideGlobal

		if extraConfig.Secrets.CustomPatternsOverrideGlobal {
			// override global custom patterns with repo custom patterns
			config.Secrets.CustomPatterns = extraConfig.Secrets.CustomPatterns
		} else if config.Secrets.CustomPatterns == nil {
			config.Secrets.CustomPatterns = extraConfig.Secrets.CustomPatterns
		} else {
			// merge repo custom patterns with global custom patterns
			for k, v := range extraConfig.Secrets.CustomPatterns {
				config.Secrets.CustomPatterns[k] = v
			}
		}

		if len(extraConfig.Notifications) > 0 && config.Notifications == nil {
			config.Notifications = extraConfig.Notifications
		}

		for k, v := range extraConfig.Notifications {
			config.Notifications[k] = v
		}

		if len(extraConfig.ScheduledNotifications) > 0 && config.ScheduledNotifications == nil {
			config.ScheduledNotifications = extraConfig.ScheduledNotifications
		}

		for k, v := range extraConfig.ScheduledNotifications {
			config.ScheduledNotifications[k] = v
		}
	}

	return &config
}
