package models

type Code struct {
	Ignore []CodeIgnore `yaml:"ignore,omitempty"`
}

type CodeIgnore struct {
	CWEs    []int    `yaml:"cwes,omitempty"`
	RuleIDs []string `yaml:"rule_ids,omitempty"`
	Reason  string   `yaml:"reason,omitempty"`
	Dirs    []string `yaml:"dirs,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`
	Expiry  string   `yaml:"expiry,omitempty"`

	// global config only
	Repositories []string `yaml:"repositories,omitempty"`
}
