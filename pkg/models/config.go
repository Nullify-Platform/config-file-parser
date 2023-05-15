package models

type Configuration struct {
	MinimumCommentSeverity string   `yaml:"minimum_comment_severity"`
	IgnoreDirs             []string `yaml:"ignore_dirs"`
	EmailNotifications     []string `yaml:"email_notifications"`
	SecretsWhitelist       []string `yaml:"secrets_whitelist"`
}
