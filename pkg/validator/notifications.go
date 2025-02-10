package validator

import (
	"fmt"
	"net/mail"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidateNotifications(config *models.Configuration) []ValidationError {
	var errors []ValidationError
	if config.Notifications == nil {
		return errors
	}

	line := 0
	column := 0

	for key, notification := range config.Notifications {
		if notification.Targets.Email == nil {
			continue
		}

		if node, exists := config.LocationInfo[fmt.Sprintf("notifications.%s.targets.email.address", key)]; exists {
			line = node.Line
			column = node.Column
		}

		if notification.Targets.Email.Address != "" {
			_, err := mail.ParseAddress(notification.Targets.Email.Address)
			if err != nil {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("notifications.%s.targets.email.address", key),
					Message: "Invalid notifications",
					Line:    line,
					Column:  column,
				})
			}
		}

		for _, email := range notification.Targets.Email.Addresses {
			_, err := mail.ParseAddress(email)
			if err != nil {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("notifications.%s.targets.email.addresses", key),
					Message: "Invalid notifications",
					Line:    line,
					Column:  column,
				})
			}
		}
	}

	return errors
}
