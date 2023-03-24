package parser

import "github.com/nullify-platform/config-file-parser/pkg/models"

const DefaultMinimumCommentSeverity = models.SeverityMedium

func NewDefaultConfig() *models.Configuration {
	return &models.Configuration{
		MinimumCommentSeverity: DefaultMinimumCommentSeverity,
		IgnoreDirs:             []string{},
	}
}
