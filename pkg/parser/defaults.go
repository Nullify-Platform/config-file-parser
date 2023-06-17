package parser

import "github.com/nullify-platform/config-file-parser/pkg/models"

const DefaultSeverityThreshold = models.SeverityMedium

func NewDefaultConfig() *models.Configuration {
	return &models.Configuration{
		SeverityThreshold: DefaultSeverityThreshold,
		IgnoreDirs:        []string{},
		IgnorePaths:       []string{},
	}
}
