package validator

import (
	"time"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateAutoFix(config *models.Configuration) bool {
	return validateAutoFix(&config.Code.AutoFix) && validateAutoFix(&config.Dependencies.AutoFix)
}

func validateAutoFix(autofix *models.AutoFix) bool {
	if !autofix.Enabled {
		return true
	}

	if autofix.MaxPullRequestsOpen < 0 {
		return false
	}

	if autofix.MaxPullRequestCreationRate == nil {
		return true
	}

	if autofix.MaxPullRequestCreationRate.Count < 0 {
		return false
	}

	_, err := time.ParseDuration(autofix.MaxPullRequestCreationRate.Period)
	return err == nil
}
