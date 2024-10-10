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
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
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
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
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
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey:        "JIRINT",
						IssueType:         "Nul-Finding",
						OnFixTransition:   "Done",
						SeverityThreshold: models.SeverityHigh,
						PriorityThreshold: models.PriorityImportant,
					},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        models.SeverityHigh,
				PriorityThreshold:        models.PriorityMedium,
				IgnoreDirs:               []string{"dir1", "dir2"},
				IgnorePaths:              []string{"path1", "path2"},
				Code: models.Code{
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
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
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey:        "JIRINT",
						IssueType:         "Nul-Finding",
						OnFixTransition:   "Done",
						SeverityThreshold: models.SeverityHigh,
						PriorityThreshold: models.PriorityImportant,
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
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
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
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey:        "JIRINT",
						IssueType:         "Nul-Finding",
						OnFixTransition:   "Done",
						SeverityThreshold: models.SeverityHigh,
						PriorityThreshold: models.PriorityImportant,
						Priorities: &models.Priorities{
							Critical: "highest",
							High:     "high",
							Medium:   "medium",
							Low:      "low",
						},
					},
				},
			},
			repoConfig: nil,
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        models.SeverityHigh,
				PriorityThreshold:        models.PriorityMedium,
				IgnoreDirs:               []string{"dir1", "dir2"},
				IgnorePaths:              []string{"path1", "path2"},
				Code: models.Code{
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
					Ignore: []models.CodeIgnore{
						{
							CWEs: []int{123},
						},
					},
				},
				Dependencies: models.Dependencies{
					AutoFix: &models.AutoFix{
						Enabled: true,
					},
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
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey:        "JIRINT",
						IssueType:         "Nul-Finding",
						OnFixTransition:   "Done",
						SeverityThreshold: models.SeverityHigh,
						PriorityThreshold: models.PriorityImportant,
						Priorities: &models.Priorities{
							Critical: "highest",
							High:     "high",
							Medium:   "medium",
							Low:      "low",
						},
					},
				},
			},
		},
		{
			name:         "repo config without severity threshold",
			globalConfig: nil,
			repoConfig: &models.Configuration{
				SeverityThreshold: "",
				PriorityThreshold: "",
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
			},
		},
		{
			name: "global config without severity threshold",
			globalConfig: &models.Configuration{
				SeverityThreshold: "",
				PriorityThreshold: "",
			},
			repoConfig: nil,
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
			},
		},
		{
			name: "global and repo config without severity threshold",
			globalConfig: &models.Configuration{
				SeverityThreshold: "",
				PriorityThreshold: "",
			},
			repoConfig: &models.Configuration{
				SeverityThreshold: "",
				PriorityThreshold: "",
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
			},
		},
		{
			name: "global and repo config without severity threshold",
			globalConfig: &models.Configuration{
				SeverityThreshold: models.SeverityCritical,
				PriorityThreshold: models.PriorityUrgent,
			},
			repoConfig: &models.Configuration{
				SeverityThreshold: models.SeverityHigh,
				PriorityThreshold: models.PriorityImportant,
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        models.SeverityHigh,
				PriorityThreshold:        models.PriorityImportant,
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			config := MergeConfigFiles(parser.NewDefaultConfig(), scenario.globalConfig, scenario.repoConfig)
			require.Equal(t, scenario.expected, config, scenario.name)
		})
	}
}

func TestMergeJira(t *testing.T) {
	for _, scenario := range []struct {
		name         string
		globalConfig *models.Configuration
		repoConfig   *models.Configuration
		expected     *models.Configuration
	}{
		{
			name: "Jira no priorities in global config, nil repoConfig",
			globalConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{},
				},
			},
			repoConfig: nil,
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "",
						IssueType:  "",
					},
				},
			},
		},
		{
			name: "Jira priorities from global config, nil repo",
			globalConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{
						Priorities: &models.Priorities{
							Critical: "urgent_global",
							High:     "high_global",
							Medium:   "medium_global",
							Low:      "low_global",
						},
						Assignee: &models.Assignee{
							Name: "John Smith",
							ID:   "123456",
						},
					},
				},
			},
			repoConfig: nil,
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "",
						IssueType:  "",
						Priorities: &models.Priorities{
							Critical: "urgent_global",
							High:     "high_global",
							Medium:   "medium_global",
							Low:      "low_global",
						},
						Assignee: &models.Assignee{
							Name: "John Smith",
							ID:   "123456",
						},
					},
				},
			},
		},
		{
			name:         "Jira priorities and assignee from repo config, nil global",
			globalConfig: nil,
			repoConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{
						Priorities: &models.Priorities{
							Critical: "urgent_repo",
							High:     "high_repo",
							Medium:   "medium_repo",
							Low:      "low_repo",
						},
						Assignee: &models.Assignee{
							Name: "John Smith",
							ID:   "123456",
						},
					},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "",
						IssueType:  "",
						Priorities: &models.Priorities{
							Critical: "urgent_repo",
							High:     "high_repo",
							Medium:   "medium_repo",
							Low:      "low_repo",
						},
						Assignee: &models.Assignee{
							Name: "John Smith",
							ID:   "123456",
						},
					},
				},
			},
		},
		{
			name: "Jira priorities from repo config overriding global config",
			globalConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{
						Priorities: &models.Priorities{
							Critical: "urgent_global",
							High:     "high_global",
							Medium:   "medium_global",
							Low:      "low_global",
						},
						Assignee: &models.Assignee{
							Name: "John Smith",
							ID:   "123456",
						},
					},
				},
			},
			repoConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{
						Priorities: &models.Priorities{
							Critical: "urgent_repo",
							High:     "high_repo",
							Medium:   "medium_repo",
							Low:      "low_repo",
						},
					},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "",
						IssueType:  "",
						Priorities: &models.Priorities{
							Critical: "urgent_repo",
							High:     "high_repo",
							Medium:   "medium_repo",
							Low:      "low_repo",
						},
						Assignee: &models.Assignee{
							Name: "John Smith",
							ID:   "123456",
						},
					},
				},
			},
		},
		{
			name:         "Jira repo config with nil priorities",
			globalConfig: nil,
			repoConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "",
						IssueType:  "",
					},
				},
			},
		},
		{
			name: "Jira repo config with project key overriding",
			globalConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "projectkey1",
						IssueType:  "issuetype1",
					},
				},
			},
			repoConfig: &models.Configuration{
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "projectkey2",
					},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Integrations: models.Integrations{
					Jira: &models.Jira{
						ProjectKey: "projectkey2",
						IssueType:  "issuetype1",
					},
				},
			},
		},
		{
			name: "secrets custom patterns global",
			globalConfig: &models.Configuration{
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1",
						},
					},
				},
			},
			repoConfig: nil,
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1",
						},
					},
				},
			},
		},
		{
			name:         "secrets custom patterns repo",
			globalConfig: nil,
			repoConfig: &models.Configuration{
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1",
						},
					},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1",
						},
					},
				},
			},
		},
		{
			name: "secrets custom patterns merge",
			globalConfig: &models.Configuration{
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1",
						},
						"custom2": {
							SecretRegex: "custom2",
						},
					},
				},
			},
			repoConfig: &models.Configuration{
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1-repo-override",
						},
						"custom3": {
							SecretRegex: "custom3",
						},
					},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1-repo-override",
						},
						"custom2": {
							SecretRegex: "custom2",
						},
						"custom3": {
							SecretRegex: "custom3",
						},
					},
				},
			},
		},
		{
			name: "secrets custom patterns override global",
			globalConfig: &models.Configuration{
				Secrets: models.Secrets{
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom1": {
							SecretRegex: "custom1",
						},
					},
				},
			},
			repoConfig: &models.Configuration{
				Secrets: models.Secrets{
					CustomPatternsOverrideGlobal: true,
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom2": {
							SecretRegex: "custom2",
						},
					},
				},
			},
			expected: &models.Configuration{
				EnablePullRequestReviews: models.Bool(true),
				EnableIssueDashboards:    models.Bool(true),
				SeverityThreshold:        parser.DefaultSeverityThreshold,
				PriorityThreshold:        parser.DefaultPriorityThreshold,
				Secrets: models.Secrets{
					CustomPatternsOverrideGlobal: true,
					CustomPatterns: map[string]models.SecretsCustomPattern{
						"custom2": {
							SecretRegex: "custom2",
						},
					},
				},
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			config := MergeConfigFiles(parser.NewDefaultConfig(), scenario.globalConfig, scenario.repoConfig)
			require.Equal(t, scenario.expected, config, scenario.name)
		})
	}
}
