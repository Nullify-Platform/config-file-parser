package models

type Secrets struct {
	Ignore []SecretsIgnore `yaml:"ignore,omitempty"`
}

type SecretsIgnore struct {
	Value   string   `yaml:"value,omitempty"`
	Pattern string   `yaml:"pattern,omitempty"`
	Reason  string   `yaml:"reason,omitempty"`
	Expiry  string   `yaml:"expiry,omitempty"`
	Dirs    []string `yaml:"dirs,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`
}
