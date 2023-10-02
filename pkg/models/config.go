package models

type Configuration struct {
	SeverityThreshold string                  `yaml:"severity_threshold"`
	IgnoreDirs        []string                `yaml:"ignore_dirs"`
	IgnorePaths       []string                `yaml:"ignore_paths"`
	SecretsWhitelist  []string                `yaml:"secrets_whitelist"` // TODO deprecate
	SecretsAllowlist  []string                `yaml:"secrets_allowlist"`
	Notifications     map[string]Notification `yaml:"notifications"`
}
