package parser

import (
	"strings"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"gopkg.in/yaml.v3"
)

func ParseConfiguration(data []byte) (*models.Configuration, error) {
	config := NewDefaultConfig()
	err := yaml.Unmarshal([]byte(data), config)
	if err != nil {
		return nil, err
	}
	sanitizeConfig(config)
	return config, nil
}

func sanitizeConfig(config *models.Configuration) {
	config.SeverityThreshold = strings.ToUpper(config.SeverityThreshold)
	if config.SeverityThreshold == "" {
		config.SeverityThreshold = DefaultSeverityThreshold
	}

	for name, n := range config.Notifications {
		n.Events.All.MinimumSeverity = strings.ToUpper(n.Events.All.MinimumSeverity)
		n.Events.NewAPIFindings.MinimumSeverity = strings.ToUpper(n.Events.NewAPIFindings.MinimumSeverity)
		n.Events.NewCodeFindings.MinimumSeverity = strings.ToUpper(n.Events.NewCodeFindings.MinimumSeverity)
		n.Events.NewDependencyFindings.MinimumSeverity = strings.ToUpper(n.Events.NewDependencyFindings.MinimumSeverity)

		config.Notifications[name] = n
	}
}
