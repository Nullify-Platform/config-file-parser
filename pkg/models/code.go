package models

type Code struct {
	EnableFailBuilds *bool        `yaml:"enable_fail_builds,omitempty"`
	Ignore           []CodeIgnore `yaml:"ignore,omitempty"`

	// TODO deprecate
	AutoFix *AutoFix `yaml:"auto_fix,omitempty"`
}

type CodeIgnore struct {
	Reason string `yaml:"reason,omitempty"`
	Expiry string `yaml:"expiry,omitempty"`

	// matchers
	CWEs    []int    `yaml:"cwes,omitempty"`
	RuleIDs []string `yaml:"rule_ids,omitempty"`
	Dirs    []string `yaml:"dirs,omitempty"`

	// global config only
	Repositories []string `yaml:"repositories,omitempty"`

	// TODO deprecate
	Paths []string `yaml:"paths,omitempty"`
}
