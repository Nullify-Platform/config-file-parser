package parser

import (
	"strings"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"gopkg.in/yaml.v3"
)

func ParseConfiguration(data []byte) (*models.Configuration, error) {
	var config models.Configuration
	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	sanitizeConfig(&config)

	return &config, nil
}

func sanitizeConfig(config *models.Configuration) {
	config.SeverityThreshold = strings.Trim(config.SeverityThreshold, " ")
	if config.SeverityThreshold != "" {
		config.SeverityThreshold = strings.ToUpper(config.SeverityThreshold)
	}

	config.PriorityThreshold = strings.Trim(config.PriorityThreshold, " ")
	if config.PriorityThreshold != "" {
		config.PriorityThreshold = strings.ToUpper(config.PriorityThreshold)
	}

	for name, n := range config.Notifications {
		if n.Events.All != nil {
			n.Events.All.MinimumSeverity = strings.ToUpper(n.Events.All.MinimumSeverity)
		}

		if n.Events.NewAPIFindings != nil {
			n.Events.NewAPIFindings.MinimumSeverity = strings.ToUpper(n.Events.NewAPIFindings.MinimumSeverity)
		}

		if n.Events.NewCodeFindings != nil {
			n.Events.NewCodeFindings.MinimumSeverity = strings.ToUpper(n.Events.NewCodeFindings.MinimumSeverity)
		}

		if n.Events.NewDependencyFindings != nil {
			n.Events.NewDependencyFindings.MinimumSeverity = strings.ToUpper(n.Events.NewDependencyFindings.MinimumSeverity)
		}

		config.Notifications[name] = n
	}
}
