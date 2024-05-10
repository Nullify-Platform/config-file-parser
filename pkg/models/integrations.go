package models

type Integrations struct {
	Jira *Jira `yaml:"jira,omitempty"`
}

type Jira struct {
	Disabled          bool        `yaml:"disabled,omitempty"`
	ProjectKey        string      `yaml:"project_key,omitempty"`
	IssueType         string      `yaml:"issue_type,omitempty"`
	SeverityThreshold string      `yaml:"severity_threshold,omitempty"`
	OnFixTransition   string      `yaml:"on_fix_transition,omitempty"`
	Priorities        *Priorities `yaml:"priorities,omitempty"`
	Assignee          *Assignee   `yaml:"assignee,omitempty"`
}

// Mapping of Nullify Finding severities to Jira Priorities.
// The user can specify the priority of the issue based on the severity.
type Priorities struct {
	Critical string `yaml:"critical,omitempty"`
	High     string `yaml:"high,omitempty"`
	Medium   string `yaml:"medium,omitempty"`
	Low      string `yaml:"low,omitempty"`
}

type Assignee struct {
	Name string `yaml:"name,omitempty"`
	ID   string `yaml:"id,omitempty"`
}
