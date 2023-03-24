package validator

import (
	"strings"
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateMinimumCommentSeverity(t *testing.T) {
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
			config:   &models.Configuration{MinimumCommentSeverity: ""},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityLow},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: strings.ToLower(models.SeverityLow)},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityMedium},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: strings.ToLower(models.SeverityMedium)},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityHigh},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: strings.ToLower(models.SeverityHigh)},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: models.SeverityCritical},
			expected: true,
		},
		{
			config:   &models.Configuration{MinimumCommentSeverity: strings.ToLower(models.SeverityCritical)},
			expected: true,
		},
		{
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
