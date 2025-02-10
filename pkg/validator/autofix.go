package validator

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateAutoFix(config *models.Configuration) []ValidationError {
	errors := []ValidationError{}
	if !validateAutoFix(config.Code.AutoFix) {
		errors = append(errors, ValidationError{
			Field:   "code.auto_fix",
			Message: "Invalid auto fix",
			Line:    config.LocationInfo["code.auto_fix"].Line,
			Column:  config.LocationInfo["code.auto_fix"].Column,
		})
	}

	if !validateAutoFix(config.Dependencies.AutoFix) {
		errors = append(errors, ValidationError{
			Field:   "dependencies.auto_fix",
			Message: "Invalid auto fix",
			Line:    config.LocationInfo["dependencies.auto_fix"].Line,
			Column:  config.LocationInfo["dependencies.auto_fix"].Column,
		})
	}
	return errors
}

func validateAutoFix(autofix *models.AutoFix) bool {
	if autofix == nil {
		return true
	}

	if !autofix.Enabled {
		return true
	}

	if autofix.MaxPullRequestsOpen != nil {
		if *autofix.MaxPullRequestsOpen < 0 {
			return false
		}
	}

	if autofix.MaxPullRequestCreationRate != nil {
		if autofix.MaxPullRequestCreationRate.Count < 0 {
			return false
		}

		if autofix.MaxPullRequestCreationRate.Days < 0 {
			return false
		}
	}

	return true
}
