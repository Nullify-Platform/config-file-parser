package validator

import (
	"strings"

	"github.com/nullify-platform/config-file-parser/pkg/models"
	"golang.org/x/exp/slices"
)

var validSeveritites = []string{
	"",
	models.SeverityLow,
	models.SeverityMedium,
	models.SeverityHigh,
	models.SeverityCritical,
	models.SeverityUnknown,
}

// ValidateMinimumCommentSeverity returns true if the minimum_comment_severity
// option is one of the valid values:
//   - ""
//   - LOW / low
//   - MEDIUM / medium
//   - HIGH / high
//   - CRITICAL / critical
//   - UNKNOWN / unknown
func ValidateMinimumCommentSeverity(config *models.Configuration) bool {
	return slices.Contains(validSeveritites, strings.ToUpper(config.MinimumCommentSeverity))
}
