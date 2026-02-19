package models

type Integrations struct {
	Jira *Jira `json:"jira,omitempty" yaml:"jira,omitempty"`
}

type Jira struct {
	Disabled            bool        `json:"disabled,omitempty"            yaml:"disabled,omitempty"`
	Enabled             *bool       `json:"enabled,omitempty"             yaml:"enabled,omitempty"`
	ProjectKey          string      `json:"projectKey,omitempty"          yaml:"project_key,omitempty"`
	IssueType           string      `json:"issueType,omitempty"           yaml:"issue_type,omitempty"`
	SeverityThreshold   string      `json:"severityThreshold,omitempty"   yaml:"severity_threshold,omitempty"`
	PriorityThreshold   string      `json:"priorityThreshold,omitempty"   yaml:"priority_threshold,omitempty"`
	OnFixTransition     string      `json:"onFixTransition,omitempty"     yaml:"on_fix_transition,omitempty"`
	CommentOnClose      *bool       `json:"commentOnClose,omitempty"      yaml:"comment_on_close,omitempty"`
	Labels              []string    `json:"labels,omitempty"              yaml:"labels,omitempty"`
	TitleTemplate       string      `json:"titleTemplate,omitempty"       yaml:"title_template,omitempty"`
	DescriptionTemplate string      `json:"descriptionTemplate,omitempty" yaml:"description_template,omitempty"`
	Priorities          *Priorities `json:"priorities,omitempty"          yaml:"priorities,omitempty"`
	Assignee            *Assignee   `json:"assignee,omitempty"            yaml:"assignee,omitempty"`
}

// Mapping of Nullify Finding severities to Jira Priorities.
// The user can specify the priority of the issue based on the severity.
type Priorities struct {
	Critical   string `json:"critical,omitempty"   yaml:"critical,omitempty"`
	High       string `json:"high,omitempty"       yaml:"high,omitempty"`
	Medium     string `json:"medium,omitempty"     yaml:"medium,omitempty"`
	Low        string `json:"low,omitempty"        yaml:"low,omitempty"`
	Urgent     string `json:"urgent,omitempty"     yaml:"urgent,omitempty"`
	Important  string `json:"important,omitempty"  yaml:"important,omitempty"`
	Negligible string `json:"negligible,omitempty" yaml:"negligible,omitempty"`
}

type Assignee struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	ID   string `json:"id,omitempty"   yaml:"id,omitempty"`
}
