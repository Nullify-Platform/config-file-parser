package validator

import (
	"path/filepath"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateGlob(config *models.Configuration) bool {
	if config.IgnorePatterns == nil {
		return true
	}

	for _, pattern := range config.IgnorePatterns {
		_, err := filepath.Glob(pattern)
		if err != nil {
			return false
		}
	}

	return true
}
