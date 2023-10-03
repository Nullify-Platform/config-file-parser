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
