package validator

import (
	"fmt"
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidGlobs(t *testing.T) {
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
			config:   &models.Configuration{IgnorePaths: []string{}},
			expected: true,
		},
		{
			config:   &models.Configuration{IgnorePaths: []string{"*[abc]"}},
			expected: true,
		},
		{
			config:   &models.Configuration{IgnorePaths: []string{"*[abc]", "*d"}},
			expected: true,
		},
		{
			config:   &models.Configuration{IgnorePaths: []string{"*[abc"}},
			expected: false,
		},
		{
			config:   &models.Configuration{IgnorePaths: []string{"*d", "*[abc"}},
			expected: false,
		},
		{
			config:   &models.Configuration{IgnorePaths: []string{"*[abc", "*d"}},
			expected: false,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			isValid := ValidateGlob(scenario.config)
			assert.Equalf(t, isValid, scenario.expected, fmt.Sprintf("failed test, globs: %s, len: %d\n", scenario.config.IgnorePaths, len(scenario.config.IgnorePaths)))
		})
	}
}

const validGlob string = `
severity_threshold: medium
ignore_dirs: ["data"]	
ignore_paths: ["*[abc]"]
email_notifications: ["hello@gmail.com"]
`
const emptyGlob string = `
severity_threshold: medium
ignore_dirs: ["data"]	
ignore_paths: ["*[abc]", "*d"]
email_notifications: ["hello@gmail.com"]
`

const twoValidGlob string = `
severity_threshold: medium
ignore_dirs: ["data"]	
ignore_paths: 
email_notifications: ["hello@gmail.com"]
`

const invalidGlob string = `
severity_threshold: medium
ignore_dirs: ["data"]	
ignore_paths: ["*[abc"]
email_notifications: ["hello@gmail.com"]
`

const endInvalidGlob string = `
severity_threshold: medium
ignore_dirs: ["data"]	
ignore_paths: ["*d", "*[abc"]
email_notifications: ["hello@gmail.com"]
`

const startInvalidGlob string = `
severity_threshold: medium
ignore_dirs: ["data"]	
ignore_paths: ["*[abc", "*d"]
email_notifications: ["hello@gmail.com"]
`

func TestParsingAndValidGlobs(t *testing.T) {
	config1, err := parser.ParseConfiguration([]byte(validGlob))
	require.NoError(t, err)
	require.Equal(t, true, ValidateGlob(config1))
	config2, err := parser.ParseConfiguration([]byte(emptyGlob))
	require.NoError(t, err)
	require.Equal(t, true, ValidateGlob(config2))
	config3, err := parser.ParseConfiguration([]byte(twoValidGlob))
	require.NoError(t, err)
	require.Equal(t, true, ValidateGlob(config3))
	config4, err := parser.ParseConfiguration([]byte(endInvalidGlob))
	require.NoError(t, err)
	require.Equal(t, false, ValidateGlob(config4))
	config5, err := parser.ParseConfiguration([]byte(startInvalidGlob))
	require.NoError(t, err)
	require.Equal(t, false, ValidateGlob(config5))
	config6, err := parser.ParseConfiguration([]byte(invalidGlob))
	require.NoError(t, err)
	require.Equal(t, false, ValidateGlob(config6))
}
