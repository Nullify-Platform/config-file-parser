package validator

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

// ValidateSeverityThreshold returns true if the severity_threshold
// option is one of the valid values:
//   - ""
//   - LOW / low
//   - MEDIUM / medium
//   - HIGH / high
//   - CRITICAL / critical
func ValidateSeverityThreshold(config *models.Configuration) []ValidationError {
	if !slices.Contains(validSeveritites, config.SeverityThreshold) {
		return []ValidationError{
			{
				Field:   "severityThreshold",
				Message: "Invalid severity threshold",
				Line:    config.LocationInfo["severity_threshold"].Line,
				Column:  config.LocationInfo["severity_threshold"].Column,
			},
		}
	}
	return []ValidationError{}
}

var validPriorities = []string{
	"",
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
func ValidatePriorityThreshold(config *models.Configuration) []ValidationError {
	if !slices.Contains(validPriorities, config.PriorityThreshold) {
		return []ValidationError{
			{
				Field:   "priorityThreshold",
				Message: "Invalid priority threshold",
				Line:    config.LocationInfo["priority_threshold"].Line,
				Column:  config.LocationInfo["priority_threshold"].Column,
			},
		}
	}
	return []ValidationError{}
}
