package parser

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configStr string = `
severity_threshold: high
ignore_dirs: ["data"]
secrets_whitelist: ["secretPassword",	 "superSecretPassword"]
`

func TestParseConfiguration(t *testing.T) {
	for _, scenario := range []struct {
		name     string
		data     string
		expected *models.Configuration
	}{
		{
			name: "default values",
			data: "",
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityMedium,
				IgnoreDirs:        nil,
				IgnorePatterns:    nil,
				SecretsWhitelist:  nil,
			},
		},
		{
			name: "user provided values",
			data: configStr,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
				IgnoreDirs:        []string{"data"},
				IgnorePatterns:    []string{"*d"},
				SecretsWhitelist:  []string{"secretPassword", "superSecretPassword"},
			},
		},
		{
			name: "user provided empty severity_threshold",
			data: "severity_threshold: ''",
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityMedium,
				IgnoreDirs:        nil,
				IgnorePatterns:    nil,
				SecretsWhitelist:  nil,
			},
		},
		{
			name: "user provided LOW severity_threshold",
			data: "severity_threshold: 'LOW'",
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityLow,
				IgnoreDirs:        nil,
				IgnorePatterns:    nil,
				SecretsWhitelist:  nil,
			},
		},
		{
			name: "user provided a single secret",
			data: `secrets_whitelist: ["password"]`,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityMedium,
				IgnoreDirs:        nil,
				IgnorePatterns:    nil,
				SecretsWhitelist:  []string{"password"},
			},
		},
		{
			name: "user provided empty secret whitelist",
			data: `secrets_whitelist: `,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityMedium,
				IgnoreDirs:        nil,
				SecretsWhitelist:  nil,
				IgnorePatterns:    nil,
			},
		},
		{
			name: "user provided empty ignore patterns",
			data: `ignore_patterns: `,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityMedium,
				IgnoreDirs:        nil,
				IgnorePatterns:    nil,
				SecretsWhitelist:  nil,
			},
		},
		{
			name: "user provided glob in ignore patterns",
			data: `ignore_patterns: ["*d"]`,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityMedium,
				IgnoreDirs:        nil,
				IgnorePatterns:    []string{"*d"},
				SecretsWhitelist:  nil,
			},
		},
		{
			name: "user provided glob in ignore patterns",
			data: `ignore_patterns: ["*d"]`,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityMedium,
				IgnoreDirs:        nil,
				IgnorePatterns:    []string{"*d"},
				SecretsWhitelist:  nil,
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			config, err := ParseConfiguration([]byte(scenario.data))
			require.NoError(t, err)
			assert.Equal(t, scenario.expected.SeverityThreshold, config.SeverityThreshold)
			require.Equal(t, len(scenario.expected.IgnoreDirs), len(config.IgnoreDirs))
			for i, v := range config.IgnoreDirs {
				assert.Equal(t, v, scenario.expected.IgnoreDirs[i])
			}
		})
	}
}
