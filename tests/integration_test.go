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
		SecretsAllowlist:  []string{"secret123"},
		Notifications: map[string]models.Notification{
			"all-events-webhook": {
				Events: models.NotificationEvents{
					All: models.NotificationEventAll{
						MinimumSeverity: models.SeverityHigh,
						SecretTypes:     []string{"ssh_key"},
					},
				},
				Targets: models.NotificationTargets{
					Webhook: models.NotificationTargetWebhook{
						URLs: []string{"https://webhook.site/123456"},
					},
				},
			},
			"findings-to-slack-and-email": {
				Events: models.NotificationEvents{
					NewCodeFindings: models.NotificationEventNewCodeFindings{
						MinimumSeverity: models.SeverityHigh,
					},
					NewSecretFindings: models.NotificationEventNewSecretFindings{
						Types: []string{"ssh_key"},
					},
					NewDependencyFindings: models.NotificationEventNewDependencyFindings{
						MinimumSeverity: models.SeverityHigh,
					},
				},
				Targets: models.NotificationTargets{
					Slack: models.NotificationTargetSlack{
						Channels: []string{"123456"},
					},
					Email: models.NotificationTargetEmail{
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
					Email: models.ScheduledNotificationTargetEmail{
						Addresses: []string{"notifications@nullify.ai", "noreply@nullify.ai"},
					},
					Slack: models.ScheduledNotificationTargetSlack{
						Channels: []string{"123456"},
					},
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
