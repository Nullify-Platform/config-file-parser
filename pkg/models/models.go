package models

type Configuration struct {
	// git platform options
	EnableFailBuilds         *bool `json:"enableFailBuilds,omitempty"         yaml:"enable_fail_builds,omitempty"`
	EnablePullRequestReviews *bool `json:"enablePullRequestReviews,omitempty" yaml:"enable_pull_request_reviews,omitempty"`
	EnableIssueDashboards    *bool `json:"enableIssueDashboards,omitempty"    yaml:"enable_issue_dashboards,omitempty"`

	SeverityThreshold string `json:"severityThreshold,omitempty" yaml:"severity_threshold,omitempty"`
	PriorityThreshold string `json:"priorityThreshold,omitempty" yaml:"priority_threshold,omitempty"`

	IgnoreDirs  []string `json:"ignoreDirs,omitempty"  yaml:"ignore_dirs,omitempty"`
	IgnorePaths []string `json:"ignorePaths,omitempty" yaml:"ignore_paths,omitempty"`

	Integrations Integrations `json:"integrations,omitempty" yaml:"integrations,omitempty"`

	// features
	Code         Code         `json:"code"         yaml:"code"`
	Dependencies Dependencies `json:"dependencies" yaml:"dependencies"`
	Secrets      Secrets      `json:"secrets"      yaml:"secrets"`

	// TODO deprecate
	SecretsWhitelist []string `json:"secretsWhitelist,omitempty" yaml:"secrets_whitelist,omitempty"`
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
