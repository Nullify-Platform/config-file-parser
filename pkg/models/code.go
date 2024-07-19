package models

type Code struct {
	FailBuilds *bool        `yaml:"fail_builds,omitempty"`
	AutoFix    *AutoFix     `yaml:"auto_fix,omitempty"`
	Ignore     []CodeIgnore `yaml:"ignore,omitempty"`
}

type CodeIgnore struct {
	Reason string `yaml:"reason,omitempty"`
	Expiry string `yaml:"expiry,omitempty"`

	// matchers
	CWEs    []int    `yaml:"cwes,omitempty"`
	RuleIDs []string `yaml:"rule_ids,omitempty"`
	Dirs    []string `yaml:"dirs,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`

	// global config only
	Repositories []string `yaml:"repositories,omitempty"`
}
