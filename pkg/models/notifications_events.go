package models

type NotificationEvents struct {
	All                   *NotificationEventAll                   `yaml:"all,omitempty"`
	NewAPIFindings        *NotificationEventNewAPIFindings        `yaml:"new_api_findings,omitempty"`
	NewCodeFindings       *NotificationEventNewCodeFindings       `yaml:"new_code_findings,omitempty"`
	NewDependencyFindings *NotificationEventNewDependencyFindings `yaml:"new_dependency_findings,omitempty"`
	NewSecretFindings     *NotificationEventNewSecretFindings     `yaml:"new_secret_findings,omitempty"`
}

type NotificationEventAll struct {
	MinimumSeverity string   `yaml:"minimum_severity,omitempty"`
	MinimumPriority int      `yaml:"minimum_priority,omitempty"`
	CWEs            []int    `yaml:"cwes,omitempty"`
	CVEs            []string `yaml:"cves,omitempty"`
	SecretTypes     []string `yaml:"secret_types,omitempty"`
}

type NotificationEventNewAPIFindings struct {
	MinimumSeverity string `yaml:"minimum_severity,omitempty"`
	MinimumPriority int    `yaml:"minimum_priority,omitempty"`
	CWEs            []int  `yaml:"cwes,omitempty"`
}

type NotificationEventNewCodeFindings struct {
	MinimumSeverity string `yaml:"minimum_severity,omitempty"`
	MinimumPriority int    `yaml:"minimum_priority,omitempty"`
	CWEs            []int  `yaml:"cwes,omitempty"`
}

type NotificationEventNewDependencyFindings struct {
	MinimumSeverity string   `yaml:"minimum_severity,omitempty"`
	MinimumPriority int      `yaml:"minimum_priority,omitempty"`
	CWEs            []int    `yaml:"cwes,omitempty"`
	CVEs            []string `yaml:"cves,omitempty"`
}

type NotificationEventNewSecretFindings struct {
	Types []string `yaml:"types,omitempty"`
}
