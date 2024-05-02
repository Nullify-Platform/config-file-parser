package models

type Integrations struct {
	Jira *Jira `yaml:"jira,omitempty"`
}

type Jira struct {
	Disabled          bool   `yaml:"disabled,omitempty"`
	ProjectKey        string `yaml:"project_key,omitempty"`
	IssueType         string `yaml:"issue_type,omitempty"`
	SeverityThreshold string `yaml:"severity_threshold,omitempty"`
}
