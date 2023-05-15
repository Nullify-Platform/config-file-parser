package validator

import "github.com/nullify-platform/config-file-parser/pkg/models"

func ValidateSecrets(config *models.Configuration) bool {
	// the list of strings can be any mix of characters so no real validation needed yet at this step
	// the yaml marshaller will check if it's not a list of strings
	if config.SecretsWhitelist == nil {
		return true
	}

	return true
}
