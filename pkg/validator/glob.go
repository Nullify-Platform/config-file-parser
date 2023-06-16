package validator

import (
	"github.com/gobwas/glob"
	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateGlob(config *models.Configuration) bool {
	if config.IgnorePatterns == nil {
		return true
	}

	for _, pattern := range config.IgnorePatterns {
		_, err := glob.Compile(pattern)
		if err != nil {
			return false
		}
	}

	return true
}
