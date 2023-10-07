package tests

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/require"
)

func TestIntegration(t *testing.T) {
	expectedConfig := &models.Configuration{
		SeverityThreshold: models.SeverityMedium,
		IgnoreDirs:        []string{"dir1"},
		IgnorePaths:       []string{"data/**/*"},
		Secrets: models.Secrets{
			Ignore: []models.SecretsIgnore{
				{
					Value:  "mocksecret123",
					Reason: "This is a test secret, it has no access to anything",
					Paths:  []string{"**/tests/*"},
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
			},
		},
		Code: models.Code{
			Ignore: []models.CodeIgnore{
				{
					CWEs:   []int{589},
					Reason: "HTTP requests with variables in tests don't matter",
					Paths:  []string{"**/tests/*"},
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
					CVE:    "CVE-2021-1234",
					Reason: "This is a false positive",
					Expiry: "2021-12-31",
				},
			},
		},
	}

	config, err := parser.LoadFromFile("nullify.yaml")
	require.NoError(t, err)

	require.Equal(t, expectedConfig, config)
}
