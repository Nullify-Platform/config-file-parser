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
	Integrations           Integrations                     `yaml:"integrations,omitempty"`

	// TODO deprecate
	SecretsWhitelist []string `yaml:"secrets_whitelist,omitempty"`
}

func (c *Configuration) GetFailBuilds() bool {
	if c.FailBuilds == nil {
		return false
	}

	return *c.FailBuilds
}

func Bool(b bool) *bool {
	return &b
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Float32(f float32) *float32 {
	return &f
}

func Float64(f float64) *float64 {
	return &f
}
