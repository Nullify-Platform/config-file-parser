package parser

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configStr string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]
`

func TestValidateMinimumCommentSeverity(t *testing.T) {
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
				IgnoreDirs:             []string{},
			},
		},
		{
			name: "",
			data: configStr,
			expected: &models.Configuration{
				MinimumCommentSeverity: models.SeverityMedium,
				IgnoreDirs:             []string{"data"},
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
