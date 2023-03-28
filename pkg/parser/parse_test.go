package parser

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configStr string = `
minimum_comment_severity: high
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
				MinimumCommentSeverity: models.SeverityMedium,
				IgnoreDirs:             nil,
			},
		},
		{
			name: "user provided values",
			data: configStr,
			expected: &models.Configuration{
				MinimumCommentSeverity: models.SeverityHigh,
				IgnoreDirs:             []string{"data"},
			},
		},
		{
			name: "user provided empty minimum_comment_severity",
			data: "minimum_comment_severity: ''",
			expected: &models.Configuration{
				MinimumCommentSeverity: models.SeverityMedium,
				IgnoreDirs:             nil,
			},
		},
		{
			name: "user provided LOW minimum_comment_severity",
			data: "minimum_comment_severity: 'LOW'",
			expected: &models.Configuration{
				MinimumCommentSeverity: models.SeverityLow,
				IgnoreDirs:             nil,
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			config, err := ParseConfiguration([]byte(scenario.data))
			require.NoError(t, err)
			assert.Equal(t, scenario.expected.MinimumCommentSeverity, config.MinimumCommentSeverity)
			require.Equal(t, len(scenario.expected.IgnoreDirs), len(config.IgnoreDirs))
			for i, v := range config.IgnoreDirs {
				assert.Equal(t, v, scenario.expected.IgnoreDirs[i])
			}
		})
	}
}
