package merger

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/require"
)

func TestMergeConfigFiles(t *testing.T) {
	for _, scenario := range []struct {
		name         string
		globalConfig *models.Configuration
		repoConfig   *models.Configuration
		expected     *models.Configuration
	}{
		{
			name:         "no config files",
			globalConfig: nil,
			repoConfig:   nil,
			expected: &models.Configuration{
				SeverityThreshold: parser.DefaultSeverityThreshold,
			},
		},
		{
			name:         "only a repo config",
			globalConfig: nil,
			repoConfig: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
				IgnoreDirs:        []string{"dir1", "dir2"},
				IgnorePaths:       []string{"path1", "path2"},
				Code: models.Code{
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					Ignore: []models.DependenciesIgnore{
						{
							CVEs: []string{"CVE-2021-1234"},
						},
					},
				},
				Secrets: models.Secrets{
					Ignore: []models.SecretsIgnore{
						{
							Value: "password",
						},
					},
				},
				Notifications: map[string]models.Notification{
					"slack": {
						Events: models.NotificationEvents{
							All: &models.NotificationEventAll{
								MinimumSeverity: models.SeverityHigh,
							},
						},
					},
				},
				ScheduledNotifications: map[string]models.ScheduledNotification{
					"slack": {
						Schedule: "0 0 * * *",
					},
				},
			},
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
				IgnoreDirs:        []string{"dir1", "dir2"},
				IgnorePaths:       []string{"path1", "path2"},
				Code: models.Code{
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					Ignore: []models.DependenciesIgnore{
						{
							CVEs: []string{"CVE-2021-1234"},
						},
					},
				},
				Secrets: models.Secrets{
					Ignore: []models.SecretsIgnore{
						{
							Value: "password",
						},
					},
				},
				Notifications: map[string]models.Notification{
					"slack": {
						Events: models.NotificationEvents{
							All: &models.NotificationEventAll{
								MinimumSeverity: models.SeverityHigh,
							},
						},
					},
				},
				ScheduledNotifications: map[string]models.ScheduledNotification{
					"slack": {
						Schedule: "0 0 * * *",
					},
				},
			},
		},
		{
			name: "only a global config",
			globalConfig: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
				IgnoreDirs:        []string{"dir1", "dir2"},
				IgnorePaths:       []string{"path1", "path2"},
				Code: models.Code{
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					Ignore: []models.DependenciesIgnore{
						{
							CVEs: []string{"CVE-2021-1234"},
						},
					},
				},
				Secrets: models.Secrets{
					Ignore: []models.SecretsIgnore{
						{
							Value: "password",
						},
					},
				},
				Notifications: map[string]models.Notification{
					"slack": {
						Events: models.NotificationEvents{
							All: &models.NotificationEventAll{
								MinimumSeverity: models.SeverityHigh,
							},
						},
					},
				},
				ScheduledNotifications: map[string]models.ScheduledNotification{
					"slack": {
						Schedule: "0 0 * * *",
					},
				},
			},
			repoConfig: nil,
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
				IgnoreDirs:        []string{"dir1", "dir2"},
				IgnorePaths:       []string{"path1", "path2"},
				Code: models.Code{
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					Ignore: []models.DependenciesIgnore{
						{
							CVEs: []string{"CVE-2021-1234"},
						},
					},
				},
				Secrets: models.Secrets{
					Ignore: []models.SecretsIgnore{
						{
							Value: "password",
						},
					},
				},
				Notifications: map[string]models.Notification{
					"slack": {
						Events: models.NotificationEvents{
							All: &models.NotificationEventAll{
								MinimumSeverity: models.SeverityHigh,
							},
						},
					},
				},
				ScheduledNotifications: map[string]models.ScheduledNotification{
					"slack": {
						Schedule: "0 0 * * *",
					},
				},
			},
		},
		{
			name:         "repo config without severity threshold",
			globalConfig: nil,
			repoConfig: &models.Configuration{
				SeverityThreshold: "",
			},
			expected: &models.Configuration{
				SeverityThreshold: parser.DefaultSeverityThreshold,
			},
		},
		{
			name: "global config without severity threshold",
			globalConfig: &models.Configuration{
				SeverityThreshold: "",
			},
			repoConfig: nil,
			expected: &models.Configuration{
				SeverityThreshold: parser.DefaultSeverityThreshold,
			},
		},
		{
			name: "global and repo config without severity threshold",
			globalConfig: &models.Configuration{
				SeverityThreshold: "",
			},
			repoConfig: &models.Configuration{
				SeverityThreshold: "",
			},
			expected: &models.Configuration{
				SeverityThreshold: parser.DefaultSeverityThreshold,
			},
		},
		{
			name: "global and repo config without severity threshold",
			globalConfig: &models.Configuration{
				SeverityThreshold: models.SeverityCritical,
			},
			repoConfig: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
			},
			expected: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			config := MergeConfigFiles(scenario.globalConfig, scenario.repoConfig)
			require.Equal(t, scenario.expected, config, scenario.name)
		})
	}
}
