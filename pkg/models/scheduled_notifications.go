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
	Schedule string                       `yaml:"schedule,omitempty"`
	Topics   ScheduledNotificationTopics  `yaml:"topics,omitempty"`
	Targets  ScheduledNotificationTargets `yaml:"targets,omitempty"`
}

type ScheduledNotificationTopics struct {
	All             bool `yaml:"all,omitempty"`
	AllNewFindings  bool `yaml:"all_new_findings,omitempty"`
	NewAPIFindings  bool `yaml:"new_api_findings,omitempty"`
	NewCodeFindings bool `yaml:"new_code_findings,omitempty"`
	NewCVEs         bool `yaml:"new_cves,omitempty"`
	NewSecrets      bool `yaml:"new_secrets,omitempty"`
	// TODO add allowlisting, fixed findings, etc
}

type ScheduledNotificationTargets struct {
	Email *ScheduledNotificationTargetEmail `yaml:"email,omitempty"`
	Slack *ScheduledNotificationTargetSlack `yaml:"slack,omitempty"`
}

type ScheduledNotificationTargetEmail struct {
	Address   string   `yaml:"address,omitempty"`
	Addresses []string `yaml:"addresses,omitempty"`
}

type ScheduledNotificationTargetSlack struct {
	Channel  string   `yaml:"channel,omitempty"`
	Channels []string `yaml:"channels,omitempty"`
}
