package models

type Configuration struct {
	SeverityThreshold string         `yaml:"severity_threshold"`
	IgnoreDirs        []string       `yaml:"ignore_dirs"`
	IgnorePaths       []string       `yaml:"ignore_paths"`
	SecretsWhitelist  []string       `yaml:"secrets_whitelist"` // TODO deprecate
	SecretsAllowlist  []string       `yaml:"secrets_allowlist"`
	Notifications     []Notification `yaml:"notifications"`
}

type Notification struct {
	Events  []NotificationEvent  `yaml:"events"`
	Targets []NotificationTarget `yaml:"targets"`
}

type NotificationEvent struct {
	Type    string                   `yaml:"type"`
	Filters NotificationEventFilters `yaml:"filters"`
}

type NotificationEventFilters struct {
	Severity string `yaml:"severity"`
}

const (
	NotificationTypeSlack   string = "slack"
	NotificationTypeWebhook string = "webhook"
	NotificationTypeEmail   string = "email"
)

type NotificationTarget struct {
	Type string `yaml:"type"`

	// optional
	WebhookSecretID string   `yaml:"webhook_secret_id"`
	WebookURL       string   `yaml:"webhook_url"`
	SlackChannel    string   `yaml:"slack_channel"`
	Emails          []string `yaml:"emails"`
}
