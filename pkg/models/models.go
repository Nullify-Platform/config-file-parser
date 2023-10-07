package models

type Configuration struct {
	SeverityThreshold      string                           `yaml:"severity_threshold,omitempty"`
	IgnoreDirs             []string                         `yaml:"ignore_dirs,omitempty"`
	IgnorePaths            []string                         `yaml:"ignore_paths,omitempty"`
	Code                   Code                             `yaml:"code,omitempty"`
	Dependencies           Dependencies                     `yaml:"dependencies,omitempty"`
	Secrets                Secrets                          `yaml:"secrets,omitempty"`
	SecretsWhitelist       []string                         `yaml:"secrets_whitelist,omitempty"` // TODO deprecate
	Notifications          map[string]Notification          `yaml:"notifications,omitempty"`
	ScheduledNotifications map[string]ScheduledNotification `yaml:"scheduled_notifications,omitempty"`
}
