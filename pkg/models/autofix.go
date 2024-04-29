package models

type AutoFix struct {
	Enabled                 bool                            `yaml:"enabled,omitempty"`
	MaxOpenPullRequests     int                             `yaml:"max_open_pull_requests,omitempty"`
	PullRequestCreationRate *AutoFixPullRequestCreationRate `yaml:"pull_request_creation_rate,omitempty"`
}

type AutoFixPullRequestCreationRate struct {
	Count  int    `yaml:"count,omitempty"`
	Period string `yaml:"period,omitempty"`
}
