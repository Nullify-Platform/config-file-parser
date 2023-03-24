package parser

import "github.com/nullify-platform/config-file-parser/pkg/models"

func GetDefaultConfig() *models.Configuration {
	return &models.Configuration{
		MinimumCommentSeverity: models.SeverityMedium,
		IgnoreDirs:             []string{},
	}
}
