package models

type AutoFix struct {
	Enabled                    bool                            `yaml:"enabled,omitempty"`
	MaxPullRequestsOpen        int                             `yaml:"max_pull_requests_open,omitempty"`
	MaxPullRequestCreationRate *AutoFixPullRequestCreationRate `yaml:"max_pull_request_creation_rate,omitempty"`
}

type AutoFixPullRequestCreationRate struct {
	Count  int    `yaml:"count,omitempty"`
	Period string `yaml:"period,omitempty"`
}
