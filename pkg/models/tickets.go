package models

type TicketCreation struct {
	Enabled bool `yaml:"enabled,omitempty"`
	Jira    Jira `yaml:"jira,omitempty"`
}

type Jira struct {
	ProjectKey string `yaml:"projectKey,omitempty"`
	IssueType  string `yaml:"issuetype,omitempty"`
}
