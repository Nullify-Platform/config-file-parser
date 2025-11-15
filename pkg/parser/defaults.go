package parser

import "github.com/nullify-platform/config-file-parser/pkg/models"

const DefaultSeverityThreshold = models.SeverityMedium
const DefaultPriorityThreshold = models.PriorityImportant

func NewDefaultConfig() *models.Configuration {
	return &models.Configuration{
		EnableFailBuilds:         nil,
		EnablePullRequestReviews: models.Bool(true),
		EnableIssueDashboards:    models.Bool(true),
		SeverityThreshold:        DefaultSeverityThreshold,
		PriorityThreshold:        DefaultPriorityThreshold,
		IgnoreDirs:               nil,
		IgnorePaths:              nil,
		Code: models.Code{
			Ignore: nil,
		},
		Dependencies: models.Dependencies{
			Ignore: nil,
		},
		Secrets: models.Secrets{
			Ignore: nil,
		},
		Integrations: models.Integrations{},
	}
}
