package models

type Integrations struct {
	Jira *Jira `yaml:"jira,omitempty"`
}

type Jira struct {
	Disabled            bool        `yaml:"disabled,omitempty"`
	Enabled             *bool       `yaml:"enabled,omitempty"`
	ProjectKey          string      `yaml:"project_key,omitempty"`
	IssueType           string      `yaml:"issue_type,omitempty"`
	SeverityThreshold   string      `yaml:"severity_threshold,omitempty"`
	PriorityThreshold   string      `yaml:"priority_threshold,omitempty"`
	OnFixTransition     string      `yaml:"on_fix_transition,omitempty"`
	CommentOnClose      *bool       `yaml:"comment_on_close,omitempty"`
	Labels              []string    `yaml:"labels,omitempty"`
	TitleTemplate       string      `yaml:"title_template,omitempty"`
	DescriptionTemplate string      `yaml:"description_template,omitempty"`
	Priorities          *Priorities `yaml:"priorities,omitempty"`
	Assignee            *Assignee   `yaml:"assignee,omitempty"`
}

// Mapping of Nullify Finding severities to Jira Priorities.
// The user can specify the priority of the issue based on the severity.
type Priorities struct {
	Critical   string `yaml:"critical,omitempty"`
	High       string `yaml:"high,omitempty"`
	Medium     string `yaml:"medium,omitempty"`
	Low        string `yaml:"low,omitempty"`
	Urgent     string `yaml:"urgent,omitempty"`
	Important  string `yaml:"important,omitempty"`
	Negligible string `yaml:"negligible,omitempty"`
}

type Assignee struct {
	Name string `yaml:"name,omitempty"`
	ID   string `yaml:"id,omitempty"`
}
