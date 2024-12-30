package validator

import (
	"context"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
)

type ValidationError struct {
	Field   string
	Message string
	Line    int
	Column  int
}

type ValidationResult struct {
	IsValid bool
	Errors  []ValidationError
}

func ValidateConfig(config *models.Configuration) ValidationResult {
	var result ValidationResult

	result.Errors = append(result.Errors, ValidateSeverityThreshold(config)...)
	result.Errors = append(result.Errors, ValidatePriorityThreshold(config)...)
	result.Errors = append(result.Errors, ValidateNotifications(config)...)
	result.Errors = append(result.Errors, ValidateScheduledNotifications(config)...)
	result.Errors = append(result.Errors, ValidatePaths(config)...)
	result.Errors = append(result.Errors, ValidateAutoFix(config)...)

	result.IsValid = len(result.Errors) == 0
	return result
}

func IsConfigValid(ctx context.Context, configString string) (ValidationResult, error) {
	parsedConfig, parseErr := parser.ParseConfiguration([]byte(configString))
	if parseErr != nil {
		// Handle YAML parsing errors
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "yaml_syntax",
				Message: parseErr.Message,
				Line:    parseErr.Line,
				Column:  parseErr.Column,
			}},
		}, nil
	}

	return ValidateConfig(parsedConfig), nil
}
