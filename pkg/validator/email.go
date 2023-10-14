package validator

import (
	"net/mail"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateEmail(config *models.Configuration) bool {
	if config.Notifications == nil {
		return true
	}

	for _, notification := range config.Notifications {
		emailTargets := notification.Targets.Email
		if emailTargets == nil {
			return true
		}

		for _, address := range emailTargets.Addresses {
			_, err := mail.ParseAddress(address)
			if err != nil {
				return false
			}
		}
	}

	return true
}
