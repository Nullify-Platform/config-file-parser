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
		Notifications: map[string]models.Notification{
			"all-events-webhook": {
				Events: models.NotificationEvents{
					All: &models.NotificationEventAll{
						MinimumSeverity: models.SeverityHigh,
						SecretTypes:     []string{"ssh_key"},
					},
				},
				Targets: models.NotificationTargets{
					Webhook: &models.NotificationTargetWebhook{
						URLs: []string{"https://webhook.site/123456"},
					},
				},
			},
			"findings-to-slack-and-email": {
				Events: models.NotificationEvents{
					NewCodeFindings: &models.NotificationEventNewCodeFindings{
						MinimumSeverity: models.SeverityHigh,
					},
					NewSecretFindings: &models.NotificationEventNewSecretFindings{
						Types: []string{"ssh_key"},
					},
					NewDependencyFindings: &models.NotificationEventNewDependencyFindings{
						MinimumSeverity: models.SeverityHigh,
					},
				},
				Targets: models.NotificationTargets{
					Slack: &models.NotificationTargetSlack{
						Channels: []string{"123456"},
					},
					Email: &models.NotificationTargetEmail{
						Addresses: []string{"notifications@nullify.ai", "noreply@nullify.ai"},
					},
				},
				Repositories: []string{
					"config-file-parser",
					"dast-action",
					"cli",
				},
			},
		},
		ScheduledNotifications: map[string]models.ScheduledNotification{
			"new-findings": {
				Schedule: "0 0 * * *",
				Topics: models.ScheduledNotificationTopics{
					All: true,
				},
				Targets: models.ScheduledNotificationTargets{
					Email: &models.ScheduledNotificationTargetEmail{
						Addresses: []string{"notifications@nullify.ai", "noreply@nullify.ai"},
					},
					Slack: &models.ScheduledNotificationTargetSlack{
						Channels: []string{"123456"},
					},
				},
				Repositories: []string{
					"config-file-parser",
					"dast-action",
					"cli",
				},
			},
		},
		Code: models.Code{
			Ignore: []models.CodeIgnore{
				{
					CWEs:   []int{589},
					Reason: "HTTP requests with variables in tests don't matter",
					Paths:  []string{"**/tests/*"},
					Repositories: []string{
						"config-file-parser",
						"dast-action",
						"cli",
					},
				},
				{
					RuleIDs: []string{"python-sql-injection"},
					Reason:  "This code won't be going live until next year but we should fix it before then",
					Expiry:  "2021-12-31",
				},
			},
		},
		Dependencies: models.Dependencies{
			Ignore: []models.DependenciesIgnore{
				{
					CVEs:   []string{"CVE-2021-1234"},
					Reason: "This is a false positive",
					Expiry: "2021-12-31",
				},
				{
					CVEs:   []string{"CVE-2021-5678"},
					Reason: "This isn't exploitable in client applications",
					Expiry: "2021-12-31",
					Repositories: []string{
						"dast-action",
						"cli",
					},
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
		AttackSurface: &models.AttackSurface{
			Enable:               true,
			EnableDNSEnumeration: true,
			DomainNames:          []string{"172.36.255.7", "example.com"},
			IgnoreDomainNames:    []string{"jira.example.com"},
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
		Notifications: map[string]models.Notification{
			"all-events-webhook": {
				Events: models.NotificationEvents{
					All: &models.NotificationEventAll{
						MinimumSeverity: models.SeverityHigh,
						SecretTypes:     []string{"ssh_key"},
					},
				},
				Targets: models.NotificationTargets{
					Webhook: &models.NotificationTargetWebhook{
						URLs: []string{"https://webhook.site/123456"},
					},
				},
			},
			"findings-to-slack-and-email": {
				Events: models.NotificationEvents{
					NewCodeFindings: &models.NotificationEventNewCodeFindings{
						MinimumSeverity: models.SeverityHigh,
					},
					NewSecretFindings: &models.NotificationEventNewSecretFindings{
						Types: []string{"ssh_key"},
					},
					NewDependencyFindings: &models.NotificationEventNewDependencyFindings{
						MinimumSeverity: models.SeverityHigh,
					},
				},
				Targets: models.NotificationTargets{
					Slack: &models.NotificationTargetSlack{
						Channels: []string{"123456"},
					},
					Email: &models.NotificationTargetEmail{
						Addresses: []string{"notifications@nullify.ai", "noreply@nullify.ai"},
					},
				},
				Repositories: []string{
					"config-file-parser",
					"dast-action",
					"cli",
				},
			},
		},
		ScheduledNotifications: map[string]models.ScheduledNotification{
			"new-findings": {
				Schedule: "0 0 * * *",
				Topics: models.ScheduledNotificationTopics{
					All: true,
				},
				Targets: models.ScheduledNotificationTargets{
					Email: &models.ScheduledNotificationTargetEmail{
						Addresses: []string{"notifications@nullify.ai", "noreply@nullify.ai"},
					},
					Slack: &models.ScheduledNotificationTargetSlack{
						Channels: []string{"123456"},
					},
				},
				Repositories: []string{
					"config-file-parser",
					"dast-action",
					"cli",
				},
			},
		},
		Code: models.Code{
			Ignore: []models.CodeIgnore{
				{
					CWEs:   []int{589},
					Reason: "HTTP requests with variables in tests don't matter",
					Paths:  []string{"**/tests/*"},
					Repositories: []string{
						"config-file-parser",
						"dast-action",
						"cli",
					},
				},
				{
					RuleIDs: []string{"python-sql-injection"},
					Reason:  "This code won't be going live until next year but we should fix it before then",
					Expiry:  "2021-12-31",
				},
			},
		},
		Dependencies: models.Dependencies{
			Ignore: []models.DependenciesIgnore{
				{
					CVEs:   []string{"CVE-2021-1234"},
					Reason: "This is a false positive",
					Expiry: "2021-12-31",
				},
				{
					CVEs:   []string{"CVE-2021-5678"},
					Reason: "This isn't exploitable in client applications",
					Expiry: "2021-12-31",
					Repositories: []string{
						"dast-action",
						"cli",
					},
				},
			},
		},
	}

	config, err := parser.LoadFromFile("empty_fail_build.yaml")
	require.NoError(t, err)

	require.Equal(t, expectedConfig, config)
	require.True(t, validator.ValidateConfig(config))
}
