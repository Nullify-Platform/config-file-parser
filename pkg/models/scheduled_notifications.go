package models

const (
	ScheduledNotificationTopicTypeAll          = "all"
	ScheduledNotificationTopicTypeCode         = "code"
	ScheduledNotificationTopicTypeIaC          = "iac"
	ScheduledNotificationTopicTypeDependencies = "dependencies"
	ScheduledNotificationTopicTypeSecrets      = "secrets"
	ScheduledNotificationTopicTypeDAST         = "dast"
)

type ScheduledNotification struct {
	Schedule string                        `yaml:"schedule,omitempty"`
	Topics   []string                      `yaml:"topics,omitempty"`
	Targets  []ScheduledNotificationTarget `yaml:"targets,omitempty"`
}

const (
	ScheduledNotificationTargetTypeEmail = "email"
	ScheduledNotificationTargetTypeSlack = "slack"
)

type ScheduledNotificationTarget struct {
	Type string `yaml:"type,omitempty"`

	// optional depending on type
	Emails    []string `yaml:"emails,omitempty"`
	ChannelID string   `yaml:"channel_id,omitempty"`
}
