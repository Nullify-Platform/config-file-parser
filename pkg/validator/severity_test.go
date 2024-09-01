package validator

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestValidateSeverityThreshold(t *testing.T) {
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
			name:     "empty SeverityThreshold",
			config:   &models.Configuration{SeverityThreshold: ""},
			expected: true,
		},
		{
			name:     "SeverityLow",
			config:   &models.Configuration{SeverityThreshold: models.SeverityLow},
			expected: true,
		},
		{
			name:     "SeverityMedium",
			config:   &models.Configuration{SeverityThreshold: models.SeverityMedium},
			expected: true,
		},
		{
			name:     "SeverityHigh",
			config:   &models.Configuration{SeverityThreshold: models.SeverityHigh},
			expected: true,
		},
		{
			name:     "SeverityCritical",
			config:   &models.Configuration{SeverityThreshold: models.SeverityCritical},
			expected: true,
		},
		{
			name:     "unexpected severity",
			config:   &models.Configuration{SeverityThreshold: "invalid-severity"},
			expected: false,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			isValid := ValidateSeverityThreshold(scenario.config)
			assert.Equal(t, scenario.expected, isValid)
		})
	}
}
