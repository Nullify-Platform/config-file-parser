package models

type Notification struct {
	Events  NotificationEvents  `yaml:"events,omitempty"`
	Targets NotificationTargets `yaml:"targets,omitempty"`

	// global config only
	Repositories []string `yaml:"repositories,omitempty"`
}
