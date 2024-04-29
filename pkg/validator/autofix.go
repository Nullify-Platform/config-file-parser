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

	if autofix.MaxOpenPullRequests < 0 {
		return false
	}

	if autofix.PullRequestCreationRate == nil {
		return true
	}

	if autofix.PullRequestCreationRate.Count < 0 {
		return false
	}

	_, err := time.ParseDuration(autofix.PullRequestCreationRate.Period)
	return err == nil
}
