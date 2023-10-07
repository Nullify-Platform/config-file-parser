package validator

import "github.com/nullify-platform/config-file-parser/pkg/models"

func ValidateSecrets(config *models.Configuration) bool {
	// the list of strings can be any mix of characters so no validation needed at this stage
	// this will change as we start accepting files
	// the parser will check if it's a valid list of strings
	if config.SecretsWhitelist == nil {
		return true
	}

	if config.Secrets.Ignore == nil {
		return true
	}

	return true
}
