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
			config:   &models.Configuration{},
			expected: true,
		},
		{
			config:   &models.Configuration{EmailNotifications: []string{}},
			expected: true,
		},
		{
			config:   &models.Configuration{EmailNotifications: []string{"john@nullify.cloud"}},
			expected: true,
		},
		{
			config:   &models.Configuration{EmailNotifications: []string{"lisa@nullify.cloud", "lisa@gmail.nullify.cloud"}},
			expected: true,
		},
		{
			config:   &models.Configuration{EmailNotifications: []string{"john@@gmail.com"}},
			expected: false,
		},
		{
			config:   &models.Configuration{EmailNotifications: []string{"john@nullify.cloud", "john@@gmail.com"}},
			expected: false,
		},
		{
			config:   &models.Configuration{EmailNotifications: []string{"helloatgmail.com"}},
			expected: false,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			isValid := ValidateEmail(scenario.config)
			assert.Equalf(t, isValid, scenario.expected, fmt.Sprintf("failed test, emails: %s, len; %d\n", scenario.config.EmailNotifications, len(scenario.config.EmailNotifications)))
		})
	}
}

const validEmail string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["hello@gmail.com"]
`

const emptyEmail string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications:
`

const noEmailConfig string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
`
const emailWithEmptyArray string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: [""]
`

const twoValidEmails string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["john@nullify.cloud", "lisa@gmail.com"]
`
const validAndInvalid string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["john()@nullify.cloud", "lisa@gmail.com"]
`

const missingCommaIncorrectQuotes string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["hello@gmail.com john@nullify.cloud"]
`

const missingComma string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["hello@gmail.com" "john@nullify.cloud"]
`

func TestParsingAndValidEmails(t *testing.T) {
	config1, err := parser.ParseConfiguration([]byte(validEmail))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config1))

	config2, err := parser.ParseConfiguration([]byte(emptyEmail))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config2))

	config3, err := parser.ParseConfiguration([]byte(noEmailConfig))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config3))

	config4, err := parser.ParseConfiguration([]byte(emailWithEmptyArray))
	require.NoError(t, err)
	require.Equal(t, false, ValidateEmail(config4))

	config5, err := parser.ParseConfiguration([]byte(twoValidEmails))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config5))

	config6, err := parser.ParseConfiguration([]byte(validAndInvalid))
	require.NoError(t, err)
	require.Equal(t, false, ValidateEmail(config6))

	config7, err := parser.ParseConfiguration([]byte(missingCommaIncorrectQuotes))
	require.NoError(t, err)
	require.Equal(t, false, ValidateEmail(config7))

	_, err = parser.ParseConfiguration([]byte(missingComma))
	require.Error(t, err)
}
