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
	config.MinimumCommentSeverity = strings.ToUpper(config.MinimumCommentSeverity)
	if config.MinimumCommentSeverity == "" {
		config.MinimumCommentSeverity = DefaultMinimumCommentSeverity
	}
}
