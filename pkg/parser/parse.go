package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"gopkg.in/yaml.v3"
)

type LocationTracker struct {
	Locations map[string]yaml.Node
}

type ParseError struct {
	Message string
	Line    int
	Column  int
}

func ParseConfiguration(data []byte) (*models.Configuration, *ParseError) {
	// Handle empty configuration case
	if len(bytes.TrimSpace(data)) == 0 {
		config := &models.Configuration{
			LocationInfo: make(map[string]yaml.Node),
		}
		sanitizeConfig(config)
		return config, nil
	}

	var config models.Configuration
	tracker := &LocationTracker{
		Locations: make(map[string]yaml.Node),
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data))

	// First, decode into a Node to preserve location information
	var node yaml.Node
	if err := decoder.Decode(&node); err != nil {
		if yamlErr, ok := err.(*yaml.TypeError); ok {
			return nil, &ParseError{
				Message: yamlErr.Errors[0],
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		return nil, &ParseError{
			Message: err.Error(),
			Line:    node.Line,
			Column:  node.Column,
		}
	}

	// recursively construct location info
	if len(node.Content) > 0 && node.Content[0].Kind == yaml.MappingNode {
		walkYAMLNode(*node.Content[0], "", tracker)
	}

	// decode into the actual configuration
	if err := node.Decode(&config); err != nil {
		return nil, &ParseError{
			Message: err.Error(),
			Line:    node.Line,
			Column:  node.Column,
		}
	}

	sanitizeConfig(&config)

	config.LocationInfo = tracker.Locations

	return &config, nil
}

func walkYAMLNode(node yaml.Node, path string, tracker *LocationTracker) {
	if node.Kind != yaml.MappingNode {
		return
	}

	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i]
		value := node.Content[i+1]

		newPath := key.Value
		if path != "" {
			newPath = path + "." + key.Value
		}

		// Store the location information
		tracker.Locations[newPath] = *value

		// Recurse into nested structures
		if value.Kind == yaml.MappingNode {
			walkYAMLNode(*value, newPath, tracker)
		}
	}
}

func sanitizeConfig(config *models.Configuration) {
	config.SeverityThreshold = strings.Trim(config.SeverityThreshold, " ")
	if config.SeverityThreshold != "" {
		config.SeverityThreshold = strings.ToUpper(config.SeverityThreshold)
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

// Error implements the error interface for ParseError
func (e *ParseError) Error() string {
	return fmt.Sprintf("yaml error at line %d, column %d: %s", e.Line, e.Column, e.Message)
}
