package models

type Configuration struct {
	// git platform options
	EnableFailBuilds         *bool `yaml:"enable_fail_builds,omitempty"`
	EnablePullRequestReviews *bool `yaml:"enable_pull_request_reviews,omitempty"`
	EnableIssueDashboards    *bool `yaml:"enable_issue_dashboards,omitempty"`

	SeverityThreshold string `yaml:"severity_threshold,omitempty"`
	PriorityThreshold string `yaml:"priority_threshold,omitempty"`

	IgnoreDirs  []string `yaml:"ignore_dirs,omitempty"`
	IgnorePaths []string `yaml:"ignore_paths,omitempty"`
	AutoFix     *AutoFix `yaml:"auto_fix,omitempty"`

	Notifications          map[string]Notification          `yaml:"notifications,omitempty"`
	ScheduledNotifications map[string]ScheduledNotification `yaml:"scheduled_notifications,omitempty"`
	Integrations           Integrations                     `yaml:"integrations,omitempty"`

	// features
	Code         Code         `yaml:"code"`
	Dependencies Dependencies `yaml:"dependencies"`
	Secrets      Secrets      `yaml:"secrets"`

	// TODO deprecate
	SecretsWhitelist []string `yaml:"secrets_whitelist,omitempty"`
}

func (c *Configuration) GetEnableFailBuilds() bool {
	if c.EnableFailBuilds == nil {
		return false
	}

	return *c.EnableFailBuilds
}

func (c *Configuration) GetEnablePullRequestReviews() bool {
	if c.EnablePullRequestReviews == nil {
		return false
	}

	return *c.EnablePullRequestReviews
}

func (c *Configuration) GetEnableIssueDashboards() bool {
	if c.EnableIssueDashboards == nil {
		return false
	}

	return *c.EnableIssueDashboards
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
