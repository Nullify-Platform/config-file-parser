package models

type Secrets struct {
	EnableFailBuilds             *bool                           `yaml:"enable_fail_builds,omitempty"`
	Ignore                       []SecretsIgnore                 `yaml:"ignore,omitempty"`
	CustomPatterns               map[string]SecretsCustomPattern `yaml:"custom_patterns,omitempty"`
	CustomPatternsOverrideGlobal bool                            `yaml:"custom_patterns_override_global,omitempty"`
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

type SecretsCustomPattern struct {
	Description      *string  `yaml:"description,omitempty"`
	SecretRegex      string   `yaml:"secret_regex,omitempty"`
	SecretRegexGroup *int     `yaml:"secret_regex_group,omitempty"`
	Entropy          *float32 `yaml:"entropy,omitempty"`
	PathRegex        *string  `yaml:"path_regex,omitempty"`
	Keywords         []string `yaml:"keywords,omitempty"`
}
