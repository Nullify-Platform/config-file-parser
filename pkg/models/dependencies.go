package models

type Dependencies struct {
	EnableFailBuilds *bool                `json:"enableFailBuilds,omitempty" yaml:"enable_fail_builds,omitempty"`
	Ignore           []DependenciesIgnore `json:"ignore,omitempty"           yaml:"ignore,omitempty"`
}

type DependenciesIgnore struct {
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Expiry string `json:"expiry,omitempty" yaml:"expiry,omitempty"`

	// matchers
	CVEs []string `json:"cves,omitempty" yaml:"cves,omitempty"`
	Dirs []string `json:"dirs,omitempty" yaml:"dirs,omitempty"`

	// global config only
	Repositories []string `json:"repositories,omitempty" yaml:"repositories,omitempty"`

	// TODO deprecate
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`
}
