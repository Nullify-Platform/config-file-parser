package models

type Configuration struct {
	SeverityThreshold  string   `yaml:"severity_threshold"`
	IgnoreDirs         []string `yaml:"ignore_dirs"`
	IgnorePaths        []string `yaml:"ignore_paths"`
	EmailNotifications []string `yaml:"email_notifications"`
	SecretsWhitelist   []string `yaml:"secrets_whitelist"`
}
