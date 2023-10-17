package validator

import (
	"github.com/gobwas/glob"
	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidatePaths(config *models.Configuration) bool {
	if config.IgnorePaths == nil {
		return true
	}

	for _, pattern := range config.IgnorePaths {
		_, err := glob.Compile(pattern)
		if err != nil {
			return false
		}
	}

	return true
}
