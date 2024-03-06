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
				SeverityThreshold: "",
				IgnoreDirs:        nil,
				IgnorePaths:       nil,
				Secrets: models.Secrets{
					Ignore: nil,
				},
			},
		},
		{
			name: "user provided values",
			data: configStr,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
				IgnoreDirs:        []string{"data"},
				IgnorePaths:       []string{"*d"},
				Secrets: models.Secrets{
					Ignore: []models.SecretsIgnore{
						{
							Value: "secretPassword",
						},
						{
							Value: "superSecretPassword",
						},
					},
				},
			},
		},
		{
			name: "user provided empty severity_threshold",
			data: "severity_threshold: ''",
			expected: &models.Configuration{
				SeverityThreshold: "",
				IgnoreDirs:        nil,
				IgnorePaths:       nil,
				Secrets: models.Secrets{
					Ignore: nil,
				},
			},
		},
		{
			name: "user provided LOW severity_threshold",
			data: "severity_threshold: 'LOW'",
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityLow,
				IgnoreDirs:        nil,
				IgnorePaths:       nil,
				Secrets: models.Secrets{
					Ignore: nil,
				},
			},
		},
		{
			name: "user provided empty ignore patterns",
			data: `ignore_paths: `,
			expected: &models.Configuration{
				SeverityThreshold: "",
				IgnoreDirs:        nil,
				IgnorePaths:       nil,
				Secrets: models.Secrets{
					Ignore: nil,
				},
			},
		},
		{
			name: "user provided glob in ignore patterns",
			data: `ignore_paths: ["*d"]`,
			expected: &models.Configuration{
				SeverityThreshold: "",
				IgnoreDirs:        nil,
				IgnorePaths:       []string{"*d"},
				Secrets: models.Secrets{
					Ignore: nil,
				},
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
