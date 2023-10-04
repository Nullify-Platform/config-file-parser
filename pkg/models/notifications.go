package models

type Notification struct {
	Events  []NotificationEvent  `yaml:"events"`
	Targets []NotificationTarget `yaml:"targets"`
}

const (
	NotificationEventTypeNewCodeFindings       string = "new-code-findings"
	NotificationEventTypeNewAPIFindings        string = "new-api-findings"
	NotificationEventTypeNewDependencyFindings string = "new-dependency-findings"
	NotificationEventTypeNewSecretFindings     string = "new-secret-findings"
)

const (
	NotificationEventGroupAlerts string = "alerts"
)

var NotificationEventGroups = map[string][]string{
	NotificationEventGroupAlerts: {
		NotificationEventTypeNewCodeFindings,
	},
}

type NotificationEvent struct {
	Type    string                   `yaml:"type"`
	Group   string                   `yaml:"group"`
	Filters NotificationEventFilters `yaml:"filters"`
}

type NotificationEventFilters struct {
	MinimumSeverity string   `yaml:"minimum_severity"`
	MinimumPriority int      `yaml:"minimum_priority"`
	CWEs            []int    `yaml:"cwes"`
	CVEs            []string `yaml:"cves"`
	SecretTypes     []string `yaml:"secret_types"`
}

const (
	NotificationTypeWebhook string = "webhook"
	NotificationTypeSlack   string = "slack"
	NotificationTypeEmail   string = "email"
)

type NotificationTarget struct {
	Type string `yaml:"type"`

	// optional fields depending on the `type`
	SecretID  string   `yaml:"secret_id"`
	URL       string   `yaml:"url"`
	ChannelID string   `yaml:"channel_id"`
	Emails    []string `yaml:"emails"`
}
