package models

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
