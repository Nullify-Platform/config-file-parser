package models

type Dependencies struct {
	EnableFailBuilds *bool                `yaml:"enable_fail_builds,omitempty"`
	Ignore           []DependenciesIgnore `yaml:"ignore,omitempty"`

	// TODO deprecate
	AutoFix *AutoFix `yaml:"auto_fix,omitempty"`
}

type DependenciesIgnore struct {
	Reason string `yaml:"reason,omitempty"`
	Expiry string `yaml:"expiry,omitempty"`

	// matchers
	CVEs []string `yaml:"cves,omitempty"`
	Dirs []string `yaml:"dirs,omitempty"`

	// global config only
	Repositories []string `yaml:"repositories,omitempty"`

	// TODO deprecate
	Paths []string `yaml:"paths,omitempty"`
}
