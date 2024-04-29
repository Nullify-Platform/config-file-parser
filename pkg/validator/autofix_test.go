package validator

import (
	"fmt"
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestValidAutoFix(t *testing.T) {
	for _, scenario := range []struct {
		name     string
		config   *models.Configuration
		expected bool
	}{
		{
			name:     "empty config",
			config:   &models.Configuration{},
			expected: true,
		},
		{
			name: "valid",
			config: &models.Configuration{
				Code: models.Code{
					AutoFix: &models.AutoFix{
						Enabled:             true,
						MaxPullRequestsOpen: models.Int(2),
						MaxPullRequestCreationRate: &models.AutoFixPullRequestCreationRate{
							Count: 1,
							Days:  1,
						},
					},
				},
				Dependencies: models.Dependencies{
					AutoFix: &models.AutoFix{
						Enabled:             true,
						MaxPullRequestsOpen: models.Int(2),
						MaxPullRequestCreationRate: &models.AutoFixPullRequestCreationRate{
							Count: 1,
							Days:  1,
						},
					},
				},
			},
			expected: true,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			isValid := ValidateAutoFix(scenario.config)
			assert.Equalf(t, isValid, scenario.expected, fmt.Sprintf("failed test: %s\n", scenario.name))
		})
	}
}
