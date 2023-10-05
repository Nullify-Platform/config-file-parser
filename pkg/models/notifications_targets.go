package models

type NotificationTargets struct {
	NotificationTargetWebhook NotificationTargetWebhook `yaml:"webhook,omitempty"`
	NotificationTargetEmail   NotificationTargetEmail   `yaml:"email,omitempty"`
	NotificationTargetSlack   NotificationTargetSlack   `yaml:"slack,omitempty"`
}

type NotificationTargetWebhook struct {
	URLs []string `yaml:"urls,omitempty"`
	URL  string   `yaml:"url,omitempty"`
}

type NotificationTargetEmail struct {
	Email  string   `yaml:"email,omitempty"`
	Emails []string `yaml:"emails,omitempty"`
}

type NotificationTargetSlack struct {
	Channel  string   `yaml:"channel,omitempty"`
	Channels []string `yaml:"channels,omitempty"`
}
