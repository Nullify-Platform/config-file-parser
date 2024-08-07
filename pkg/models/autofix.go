package models

type AutoFix struct {
	Enabled                    bool                            `yaml:"enabled,omitempty"`
	MaxPullRequestsOpen        *int                            `yaml:"max_pull_requests_open,omitempty"`
	MaxPullRequestCreationRate *AutoFixPullRequestCreationRate `yaml:"max_pull_request_creation_rate,omitempty"`
	Labels                     []string                        `yaml:"labels,omitempty"`
}

type AutoFixPullRequestCreationRate struct {
	Count int `yaml:"count,omitempty"`
	Days  int `yaml:"days,omitempty"`
}
