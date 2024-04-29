package models

type Configuration struct {
	FailBuilds        *bool    `yaml:"fail_builds,omitempty"`
	SeverityThreshold string   `yaml:"severity_threshold,omitempty"`
	IgnoreDirs        []string `yaml:"ignore_dirs,omitempty"`
	IgnorePaths       []string `yaml:"ignore_paths,omitempty"`

	Code         Code         `yaml:"code,omitempty"`
	Dependencies Dependencies `yaml:"dependencies,omitempty"`
	Secrets      Secrets      `yaml:"secrets,omitempty"`

	Notifications          map[string]Notification          `yaml:"notifications,omitempty"`
	ScheduledNotifications map[string]ScheduledNotification `yaml:"scheduled_notifications,omitempty"`

	// TODO deprecate
	SecretsWhitelist []string `yaml:"secrets_whitelist,omitempty"`
}

func (c *Configuration) GetFailBuilds() bool {
	if c.FailBuilds == nil {
		return false
	}

	return *c.FailBuilds
}

func Int(i int) *int {
	return &i
}
