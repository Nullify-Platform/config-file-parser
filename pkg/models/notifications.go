package models

type Notification struct {
	Events  NotificationEvents  `yaml:"events,omitempty"`
	Targets NotificationTargets `yaml:"targets,omitempty"`
}
