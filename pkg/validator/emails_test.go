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
			fmt.Printf("testing config: %v\n", scenario.config.EmailNotifications)
			assert.Equalf(t, isValid, scenario.expected, fmt.Sprintf("failed test, emails: %s, len; %d\n", scenario.config.EmailNotifications, len(scenario.config.EmailNotifications)))
		})
	}
}

const configStr1 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["hello@gmail.com"]
`

const configStr2 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications:
`

const configStr3 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
`
const configStr4 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: [""]
`

const configStr5 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["john@nullify.cloud", "lisa@gmail.com"]
`
const configStr6 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["john()@nullify.cloud", "lisa@gmail.com"]
`

const configStr7 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["hello@gmail.com john@nullify.cloud"]
`

const configStr8 string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]	
email_notifications: ["hello@gmail.com" "john@nullify.cloud"]
`

func TestParsingAndValidEmails(t *testing.T) {
	config1, err := parser.ParseConfiguration([]byte(configStr1))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config1))

	config2, err := parser.ParseConfiguration([]byte(configStr2))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config2))

	config3, err := parser.ParseConfiguration([]byte(configStr3))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config3))

	config4, err := parser.ParseConfiguration([]byte(configStr4))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config4))

	config5, err := parser.ParseConfiguration([]byte(configStr5))
	require.NoError(t, err)
	require.Equal(t, true, ValidateEmail(config5))

	config6, err := parser.ParseConfiguration([]byte(configStr6))
	require.NoError(t, err)
	require.Equal(t, false, ValidateEmail(config6))

	config7, err := parser.ParseConfiguration([]byte(configStr7))
	require.NoError(t, err)
	require.Equal(t, false, ValidateEmail(config7))

	_, err = parser.ParseConfiguration([]byte(configStr8))
	require.Error(t, err)
}
