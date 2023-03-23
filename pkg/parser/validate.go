package parser

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
	"golang.org/x/exp/slices"
)

var validSeveritites = []string{
	"",
	models.SeverityLow,
	models.SeverityMedium,
	models.SeverityHigh,
	models.SeverityCritical,
}

func ValidateConfig(config *models.Configuration) bool {
	return slices.Contains(validSeveritites, config.MinimumCommentSeverity)
}
