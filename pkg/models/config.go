package models

type Configuration struct {
	SeverityThreshold  string   `yaml:"severity_threshold"`
	IgnoreDirs         []string `yaml:"ignore_dirs"`
	EmailNotifications []string `yaml:"email_notifications"`
	SecretsWhitelist   []string `yaml:"secrets_whitelist"`
}
