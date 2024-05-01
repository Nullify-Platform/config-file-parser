package models

type Jira struct {
	Project   string `yaml:"project,omitempty"`
	IssueType string `yaml:"issuetype,omitempty"`
}
