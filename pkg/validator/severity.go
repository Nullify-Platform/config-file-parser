package validator

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
	"golang.org/x/exp/slices"
)

var validSeveritites = []string{
	models.SeverityLow,
	models.SeverityMedium,
	models.SeverityHigh,
	models.SeverityCritical,
}

// ValidateSeverityThreshold returns true if the severity_threshold
// option is one of the valid values:
//   - ""
//   - LOW / low
//   - MEDIUM / medium
//   - HIGH / high
//   - CRITICAL / critical
func ValidateSeverityThreshold(config *models.Configuration) bool {
	return slices.Contains(validSeveritites, config.SeverityThreshold)
}

var validPriorities = []string{
	models.PriorityUrgent,
	models.PriorityImportant,
	models.PriorityMedium,
	models.PriorityLow,
	models.PriorityNegligible,
}

// ValidatePriorityThreshold returns true if the priority_threshold
// option is one of the valid values:
//   - ""
//   - NEGLIGIBLE / negligible
//   - LOW / low
//   - MEDIUM / medium
//   - IMPORTANT / important
//   - URGENT / urgent
func ValidatePriorityThreshold(config *models.Configuration) bool {
	return slices.Contains(validPriorities, config.PriorityThreshold)
}
