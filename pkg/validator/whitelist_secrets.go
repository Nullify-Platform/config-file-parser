package validator

import "github.com/nullify-platform/config-file-parser/pkg/models"

func ValidateSecrets(config *models.Configuration) bool {
	if config.SecretsWhitelist == nil {
		return true
	}

	return true
}
