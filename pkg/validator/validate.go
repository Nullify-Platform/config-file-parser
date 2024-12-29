package validator

import (
	"context"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/pkg/errors"
)

// ValidateConfig return true if provided configuration is valid
func ValidateConfig(config *models.Configuration) bool {
	return ValidateSeverityThreshold(config) &&
		ValidateNotifications(config) &&
		ValidateScheduledNotifications(config) &&
		ValidatePaths(config) &&
		ValidateAutoFix(config)
}

func IsConfigValid(ctx context.Context, configString string) (bool, error) {
	parsedConfig, err := parser.ParseConfiguration([]byte(configString))
	if err != nil {
		return false, errors.Wrap(err, "failed to parse config")
	}

	isValid := ValidateConfig(parsedConfig)
	return isValid, nil
}
