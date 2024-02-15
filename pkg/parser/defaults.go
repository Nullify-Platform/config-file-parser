package parser

import "github.com/nullify-platform/config-file-parser/pkg/models"

const DefaultSeverityThreshold = models.SeverityMedium

func NewDefaultConfig() *models.Configuration {
	return &models.Configuration{
		FailBuilds:             nil,
		SeverityThreshold:      DefaultSeverityThreshold,
		IgnoreDirs:             []string{},
		IgnorePaths:            []string{},
		Code:                   models.Code{},
		Dependencies:           models.Dependencies{},
		Secrets:                models.Secrets{},
		SecretsWhitelist:       []string{},
		Notifications:          map[string]models.Notification{},
		ScheduledNotifications: map[string]models.ScheduledNotification{},
	}
}
