package models

type Dependencies struct {
	Ignore []DependenciesIgnore `yaml:"ignore,omitempty"`
}

type DependenciesIgnore struct {
	CVE    string `yaml:"cve,omitempty"`
	Reason string `yaml:"reason,omitempty"`
	Expiry string `yaml:"expiry,omitempty"`
}
