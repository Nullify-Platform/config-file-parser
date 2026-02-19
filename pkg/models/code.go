package models

type Code struct {
	EnableFailBuilds *bool        `json:"enableFailBuilds,omitempty" yaml:"enable_fail_builds,omitempty"`
	Ignore           []CodeIgnore `json:"ignore,omitempty"           yaml:"ignore,omitempty"`
}

type CodeIgnore struct {
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Expiry string `json:"expiry,omitempty" yaml:"expiry,omitempty"`

	// matchers
	CWEs    []int    `json:"cwes,omitempty"    yaml:"cwes,omitempty"`
	RuleIDs []string `json:"ruleIds,omitempty" yaml:"rule_ids,omitempty"`
	Dirs    []string `json:"dirs,omitempty"    yaml:"dirs,omitempty"`

	// global config only
	Repositories []string `json:"repositories,omitempty" yaml:"repositories,omitempty"`

	// TODO deprecate
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`
}
