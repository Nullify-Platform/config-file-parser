package validator

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateAutoFix(config *models.Configuration) bool {
	return validateAutoFix(config.Code.AutoFix) && validateAutoFix(config.Dependencies.AutoFix)
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
