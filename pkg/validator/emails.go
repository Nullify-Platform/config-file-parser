package validator

import (
	"net/mail"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateEmail(config *models.Configuration) bool {
	if config.EmailNotifications == nil {
		return true
	}

	if len(config.EmailNotifications) == 1 && config.EmailNotifications[0] == "" {
		return true
	}

	for _, email := range config.EmailNotifications {
		_, err := mail.ParseAddress(email)
		if err != nil {
			return false
		}
	}

	return true
}
