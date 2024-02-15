package parser

import "github.com/nullify-platform/config-file-parser/pkg/models"

func NewDefaultConfig() *models.Configuration {
	return &models.Configuration{
		SeverityThreshold:      "",
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
