package validator

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestValidateMinimumCommentSeverity(t *testing.T) {
	for _, scenario := range []struct {
		name     string
		config   *models.Configuration
		expected bool
	}{
		{
			name:     "empty configuration",
			config:   &models.Configuration{},
			expected: false,
		},
		{
			name:     "default configuration",
			config:   parser.NewDefaultConfig(),
			expected: true,
		},
		{
			name:     "empty MinimumCommentSeverity",
			config:   &models.Configuration{MinimumCommentSeverity: ""},
			expected: false,
		},
		{
			name:     "SeverityLow",
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityLow},
			expected: true,
		},
		{
			name:     "SeverityMedium",
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityMedium},
			expected: true,
		},
		{
			name:     "SeverityHigh",
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityHigh},
			expected: true,
		},
		{
			name:     "SeverityCritical",
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityCritical},
			expected: true,
		},
		{
			name:     "unexpected severity",
			config:   &models.Configuration{MinimumCommentSeverity: "invalid-severity"},
			expected: false,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			isValid := ValidateMinimumCommentSeverity(scenario.config)
			assert.Equal(t, scenario.expected, isValid)
		})
	}
}
