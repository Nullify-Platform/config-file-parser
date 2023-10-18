package validator

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/robfig/cron/v3"
)

func ValidateScheduledNotifications(config *models.Configuration) bool {
	if config.ScheduledNotifications == nil {
		return true
	}

	for _, notification := range config.ScheduledNotifications {
		if !validateScheduledNotificationSchedule(notification.Schedule, notification.Timezone) {
			return false
		}

		if !validateScheduledNotificationEmails(notification) {
			return false
		}
	}

	return true
}

// validateScheduledNotificationSchedule return true if provided schedule is a valid cron expression.
// The cron expression can also only trigger at most once per hour.
func validateScheduledNotificationSchedule(schedule string, timezone string) bool {
	spec := "TZ=" + timezone + " " + schedule

	if timezone == "" {
		spec = "TZ=UTC " + schedule
	}

	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)

	// TODO this function panics with the following input "TZ=UTC"
	cronSchedule, err := p.Parse(spec)
	if err != nil {
		fmt.Printf("failed to parse cron expression: %s", err.Error())
		return false
	}

	// check if the cron expression triggers more often than once per hour
	start := cronSchedule.Next(time.Now())
	finish := cronSchedule.Next(start)

	return finish.Sub(start) >= time.Hour
}

func validateScheduledNotificationEmails(notification models.ScheduledNotification) bool {
	if notification.Targets.Email == nil {
		return true
	}

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

	return true
}
