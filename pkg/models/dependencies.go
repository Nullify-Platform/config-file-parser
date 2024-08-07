package models

type Dependencies struct {
	FailBuilds *bool                `yaml:"fail_builds,omitempty"`
	AutoFix    *AutoFix             `yaml:"auto_fix,omitempty"`
	Ignore     []DependenciesIgnore `yaml:"ignore,omitempty"`
}

type DependenciesIgnore struct {
	Reason string `yaml:"reason,omitempty"`
	Expiry string `yaml:"expiry,omitempty"`

	// matchers
	CVEs  []string `yaml:"cves,omitempty"`
	Dirs  []string `yaml:"dirs,omitempty"`
	Paths []string `yaml:"paths,omitempty"`

	// global config only
	Repositories []string `yaml:"repositories,omitempty"`
}
