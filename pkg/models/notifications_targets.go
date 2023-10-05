package models

type NotificationTargets struct {
	Webhook NotificationTargetWebhook `yaml:"webhook,omitempty"`
	Email   NotificationTargetEmail   `yaml:"email,omitempty"`
	Slack   NotificationTargetSlack   `yaml:"slack,omitempty"`
}

type NotificationTargetWebhook struct {
	URLs []string `yaml:"urls,omitempty"`
	URL  string   `yaml:"url,omitempty"`
}

type NotificationTargetEmail struct {
	Address   string   `yaml:"address,omitempty"`
	Addresses []string `yaml:"addresses,omitempty"`
}

type NotificationTargetSlack struct {
	Channel  string   `yaml:"channel,omitempty"`
	Channels []string `yaml:"channels,omitempty"`
}
