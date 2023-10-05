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
		if notification.Targets.NotificationTargetEmail.Email != "" {
			_, err := mail.ParseAddress(notification.Targets.NotificationTargetEmail.Email)
			if err != nil {
				return false
			}
		}

		for _, email := range notification.Targets.NotificationTargetEmail.Emails {
			_, err := mail.ParseAddress(email)
			if err != nil {
				return false
			}
		}
	}

	return true
}
