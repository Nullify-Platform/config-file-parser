package validator

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
)

// ValidateConfig return true if provided configuration is valid
func ValidateConfig(config *models.Configuration) bool {
	return ValidateSeverityThreshold(config) &&
		ValidateNotifications(config) &&
		ValidateScheduledNotifications(config) &&
		ValidatePaths(config) &&
		ValidateAutoFix(config)
}
