package models

type Secrets struct {
	EnableFailBuilds             *bool                           `json:"enableFailBuilds,omitempty"             yaml:"enable_fail_builds,omitempty"`
	Ignore                       []SecretsIgnore                 `json:"ignore,omitempty"                       yaml:"ignore,omitempty"`
	CustomPatterns               map[string]SecretsCustomPattern `json:"customPatterns,omitempty"                yaml:"custom_patterns,omitempty"`
	CustomPatternsOverrideGlobal bool                            `json:"customPatternsOverrideGlobal,omitempty" yaml:"custom_patterns_override_global,omitempty"`
}

type SecretsIgnore struct {
	Reason string `json:"reason,omitempty"  yaml:"reason,omitempty"`
	Expiry string `json:"expiry,omitempty"  yaml:"expiry,omitempty"`

	// matchers
	Value   string `json:"value,omitempty"   yaml:"value,omitempty"`
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	SHA256  string `json:"sha256,omitempty"  yaml:"sha256,omitempty"`

	// global config only
	Repositories []string `json:"repositories,omitempty" yaml:"repositories,omitempty"`

	// TODO deprecate
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`
}

type SecretsCustomPattern struct {
	Description      *string  `json:"description,omitempty"      yaml:"description,omitempty"`
	SecretRegex      string   `json:"secretRegex,omitempty"      yaml:"secret_regex,omitempty"`
	SecretRegexGroup *int     `json:"secretRegexGroup,omitempty" yaml:"secret_regex_group,omitempty"`
	Entropy          *float32 `json:"entropy,omitempty"          yaml:"entropy,omitempty"`
	PathRegex        *string  `json:"pathRegex,omitempty"        yaml:"path_regex,omitempty"`
	Keywords         []string `json:"keywords,omitempty"         yaml:"keywords,omitempty"`
}
