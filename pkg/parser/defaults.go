package parser

import "github.com/nullify-platform/config-file-parser/pkg/models"

const DefaultSeverityThreshold = models.SeverityMedium

func NewDefaultConfig() *models.Configuration {
	return &models.Configuration{
		FailBuilds:        nil,
		SeverityThreshold: DefaultSeverityThreshold,
		IgnoreDirs:        nil,
		IgnorePaths:       nil,
		Code: models.Code{
			Ignore: nil,
		},
		Dependencies: models.Dependencies{
			Ignore: nil,
		},
		Secrets: models.Secrets{
			Ignore: nil,
		},
		Notifications:          nil,
		ScheduledNotifications: nil,
	}
}
