package validator

import (
	"fmt"
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"

	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidEmails(t *testing.T) {
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
			name: "empty notifications",
			config: &models.Configuration{
				Notifications: map[string]models.Notification{
					"test": {
						Targets: models.NotificationTargets{},
					},
				},
			},
			expected: true,
		},
		{
			name: "single correct email",
			config: &models.Configuration{
				Notifications: map[string]models.Notification{
					"test": {
						Targets: models.NotificationTargets{
							Email: &models.NotificationTargetEmail{
								Addresses: []string{"john@nullify.ai"},
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "two correct emails",
			config: &models.Configuration{
				Notifications: map[string]models.Notification{
					"test": {
						Targets: models.NotificationTargets{
							Email: &models.NotificationTargetEmail{
								Addresses: []string{"lisa@nullify.ai", "lisa@gmail.nullify.ai"},
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "single incorrect email",
			config: &models.Configuration{
				Notifications: map[string]models.Notification{
					"test": {
						Targets: models.NotificationTargets{
							Email: &models.NotificationTargetEmail{
								Addresses: []string{"john@@gmail.com"},
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "one correct and one incorrect email",
			config: &models.Configuration{
				Notifications: map[string]models.Notification{
					"test": {
						Targets: models.NotificationTargets{
							Email: &models.NotificationTargetEmail{
								Addresses: []string{"john@nullify.ai", "john@@gmail.com"},
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "one incorrect email",
			config: &models.Configuration{
				Notifications: map[string]models.Notification{
					"test": {
						Targets: models.NotificationTargets{
							Email: &models.NotificationTargetEmail{
								Addresses: []string{"helloatgmail.com"},
							},
						},
					},
				},
			},
			expected: false,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			isValid := ValidateNotifications(scenario.config)
			assert.Equalf(t, isValid, scenario.expected, fmt.Sprintf("failed test: %s\n", scenario.name))
		})
	}
}

const validEmail string = `
notifications:
  test:
    targets:
      email:
        addresses: ["hello@gmail.com"]
`

const emptyEmail string = `
notifications:
  test:
    targets:
      email:
        addresses: 
`

const noEmailConfig string = `
notifications:
`

const emailWithEmptyArray string = `
notifications:
  test:
    targets:
      email:
        addresses: [""]
`

const twoValidEmails string = `
notifications:
  test:
    targets:
      email:
        addresses: ["john@nullify.ai", "lisa@gmail.com"]
`
const validAndInvalid string = `
notifications:
  test:
    targets:
      email:
        addresses: ["john()@nullify.ai", "lisa@gmail.com"]
`

const missingCommaIncorrectQuotes string = `
notifications:
  test:
    targets:
      email:
        addresses: ["hello@gmail.com john@nullify.ai"]
`

const missingComma string = `
notifications:
  test:
    targets:
      email:
        addresses: ["hello@gmail.com" "john@nullify.ai"]
`

func TestParsingAndValidEmails(t *testing.T) {
	config1, err := parser.ParseConfiguration([]byte(validEmail))
	require.NoError(t, err)
	require.Equal(t, true, ValidateNotifications(config1))

	config2, err := parser.ParseConfiguration([]byte(emptyEmail))
	require.NoError(t, err)
	require.Equal(t, true, ValidateNotifications(config2))

	config3, err := parser.ParseConfiguration([]byte(noEmailConfig))
	require.NoError(t, err)
	require.Equal(t, true, ValidateNotifications(config3))

	config4, err := parser.ParseConfiguration([]byte(emailWithEmptyArray))
	require.NoError(t, err)
	require.Equal(t, false, ValidateNotifications(config4))

	config5, err := parser.ParseConfiguration([]byte(twoValidEmails))
	require.NoError(t, err)
	require.Equal(t, true, ValidateNotifications(config5))

	config6, err := parser.ParseConfiguration([]byte(validAndInvalid))
	require.NoError(t, err)
	require.Equal(t, false, ValidateNotifications(config6))

	config7, err := parser.ParseConfiguration([]byte(missingCommaIncorrectQuotes))
	require.NoError(t, err)
	require.Equal(t, false, ValidateNotifications(config7))

	_, err = parser.ParseConfiguration([]byte(missingComma))
	require.Error(t, err)
}
