package models

type Secrets struct {
	Ignore []SecretsIgnore `yaml:"ignore,omitempty"`
}

type SecretsIgnore struct {
	Reason string `yaml:"reason,omitempty"`
	Expiry string `yaml:"expiry,omitempty"`

	// matchers
	Value   string `yaml:"value,omitempty"`
	Pattern string `yaml:"pattern,omitempty"`
	SHA256  string `yaml:"sha256,omitempty"`

	// global config only
	Repositories []string `yaml:"repositories,omitempty"`
}
