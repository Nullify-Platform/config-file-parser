package models

type Configuration struct {
	SeverityThreshold string                  `yaml:"severity_threshold,omitempty"`
	IgnoreDirs        []string                `yaml:"ignore_dirs,omitempty"`
	IgnorePaths       []string                `yaml:"ignore_paths,omitempty"`
	SecretsWhitelist  []string                `yaml:"secrets_whitelist,omitempty"` // TODO deprecate
	SecretsAllowlist  []string                `yaml:"secrets_allowlist,omitempty"`
	Notifications     map[string]Notification `yaml:"notifications,omitempty"`
}
