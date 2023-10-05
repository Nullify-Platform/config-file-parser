package validator

import (
	"net/mail"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateNotifications(config *models.Configuration) bool {
	if config.Notifications == nil {
		return true
	}

	for _, notification := range config.Notifications {
		if notification.Targets.Email.Address != "" {
			_, err := mail.ParseAddress(notification.Targets.Email.Address)
			if err != nil {
				return false
			}
		}

		for _, email := range notification.Targets.Email.Addresses {
			_, err := mail.ParseAddress(email)
			if err != nil {
				return false
			}
		}
	}

	for _, notification := range config.ScheduledNotifications {
		if notification.Targets.Email.Address != "" {
			_, err := mail.ParseAddress(notification.Targets.Email.Address)
			if err != nil {
				return false
			}
		}

		for _, email := range notification.Targets.Email.Addresses {
			_, err := mail.ParseAddress(email)
			if err != nil {
				return false
			}
		}
	}

	return true
}
