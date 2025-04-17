package validator

import (
	"fmt"
	"testing"

	"github.com/nullify-platform/config-file-parser/v2/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateScheduledNotifications(t *testing.T) {
	for _, scenario := range []struct {
		name     string
		config   *models.Configuration
		expected bool
	}{
		{
			name:     "empty config",
			config:   &models.Configuration{},
			expected: true,
		},
		{
			name: "empty scheduled notifications",
			config: &models.Configuration{
				ScheduledNotifications: map[string]models.ScheduledNotification{},
			},
			expected: true,
		},
		{
			name: "cron expression triggers every 59 minutes",
			config: &models.Configuration{
				ScheduledNotifications: map[string]models.ScheduledNotification{
					"test": {
						Schedule: "*/59 * * * *",
					},
				},
			},
			expected: false,
		},
		{
			name: "cron expression triggers every hour",
			config: &models.Configuration{
				ScheduledNotifications: map[string]models.ScheduledNotification{
					"test": {
						Schedule: "0 * * * *",
					},
				},
			},
			expected: true,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			isValid := ValidateScheduledNotifications(scenario.config)
			assert.Equalf(t, scenario.expected, isValid, fmt.Sprintf("failed test: %s\n", scenario.name))
		})
	}
}
