package models

type AutoFix struct {
	Enabled                    bool                            `yaml:"enabled,omitempty"`
	MaxPullRequestsOpen        *int                            `yaml:"max_pull_requests_open,omitempty"`
	MaxPullRequestCreationRate *AutoFixPullRequestCreationRate `yaml:"max_pull_request_creation_rate,omitempty"`
	Scripts                    *AutoFixScripts                 `yaml:"scripts,omitempty"`
}

type AutoFixPullRequestCreationRate struct {
	Count int `yaml:"count,omitempty"`
	Days  int `yaml:"days,omitempty"`
}

type AutoFixScripts struct {
	// docker image - run locally if empty
	Image   string `yaml:"image,omitempty"`
	Lint    string `yaml:"lint,omitempty"`
	Compile string `yaml:"compile,omitempty"`
	Test    string `yaml:"test,omitempty"`
}
