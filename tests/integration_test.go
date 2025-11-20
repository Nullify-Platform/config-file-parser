package tests

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/nullify-platform/config-file-parser/pkg/validator"
	"github.com/stretchr/testify/require"
)

func TestIntegration(t *testing.T) {
	expectedConfig := &models.Configuration{
		EnableFailBuilds:         models.Bool(true),
		EnablePullRequestReviews: models.Bool(true),
		EnableIssueDashboards:    models.Bool(true),
		SeverityThreshold:        models.SeverityMedium,
		PriorityThreshold:        models.PriorityUrgent,
		IgnoreDirs:               []string{"dir1"},
		IgnorePaths:              []string{"data/**/*"},
		Secrets: models.Secrets{
			Ignore: []models.SecretsIgnore{
				{
					Value:  "mocksecret123",
					Reason: "This is a test secret, it has no access to anything",
				},
				{
					Pattern: "id[0-9]+",
					Reason:  "These are not secrets, they are internal identifiers",
				},
				{
					Value:  "actualsecret123",
					Reason: "We can't remove this right now but we should",
					Expiry: "2021-12-31",
				},
			},
			CustomPatterns: map[string]models.SecretsCustomPattern{
				"my-custom-rule-1": {
					Description:      models.String("This is a custom rule for finding secrets"),
					SecretRegex:      "mysecret[0-9]+",
					SecretRegexGroup: models.Int(0),
					Entropy:          models.Float32(4.0),
					PathRegex:        models.String(".*"),
					Keywords:         []string{"package", "func"},
				},
			},
		},
		Integrations: models.Integrations{
			Jira: &models.Jira{
				Disabled:          false,
				ProjectKey:        "JIRINT",
				IssueType:         "Nul-Finding",
				SeverityThreshold: models.SeverityHigh,
				PriorityThreshold: models.PriorityImportant,
				OnFixTransition:   "Done",
			},
		},
	}

	config, err := parser.LoadFromFile("nullify.yaml")
	require.NoError(t, err)

	require.Equal(t, expectedConfig, config)
	require.True(t, validator.ValidateConfig(config))
}

func TestEmptyFailsBuildField(t *testing.T) {
	expectedConfig := &models.Configuration{
		EnableFailBuilds:  nil,
		SeverityThreshold: models.SeverityMedium,
		IgnoreDirs:        []string{"dir1"},
		IgnorePaths:       []string{"data/**/*"},
		Secrets: models.Secrets{
			Ignore: []models.SecretsIgnore{
				{
					Value:  "mocksecret123",
					Reason: "This is a test secret, it has no access to anything",
				},
				{
					Pattern: "id[0-9]+",
					Reason:  "These are not secrets, they are internal identifiers",
				},
				{
					Value:  "actualsecret123",
					Reason: "We can't remove this right now but we should",
					Expiry: "2021-12-31",
				},
			},
		},
	}

	config, err := parser.LoadFromFile("empty_fail_build.yaml")
	require.NoError(t, err)

	require.Equal(t, expectedConfig, config)
	require.True(t, validator.ValidateConfig(config))
}
