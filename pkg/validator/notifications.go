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
		for _, target := range notification.Targets {
			for _, email := range target.Emails {
				_, err := mail.ParseAddress(email)
				if err != nil {
					return false
				}
			}
		}
	}

	return true
}
