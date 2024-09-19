package validator

import (
	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateProjects(config *models.Configuration) bool {
	if config.Projects == nil {
		return true
	}

	for _, project := range config.Projects {
		if project.Path == "" {
			return false
		}
	}

	return true
}
